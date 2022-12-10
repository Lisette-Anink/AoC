module Day9.Main exposing (..)

import Browser
import Element exposing (Element, el, text, row,  column, alignRight, fill, width, rgb255, spacing, centerY, padding)
-- import Html exposing (Html, button, div, text, p)
import Element.Background as Background
import Element.Border as Border
import Element.Font as Font
import Html.Events exposing (onClick)
import Dict
import List exposing (..)
import Set

main =
  Browser.sandbox { init = init, update = update, view = view }

-- type alias Model

init =
  0


type Msg = Start

update msg model =
  case msg of
    Start ->
      model


view model =
  Element.layout []
      <| column [spacing 10]
          <|
          -- part1 data1
            -- ++ part1 dataR
           part2 data1
          ++ part2 data2
          ++ part2 dataR

part1 input =
  [el [] (text "sol part 1: ")
  , el [] (text <| String.fromInt <| findPositions <| processInput input)
   ]

part2 input =
  [el [] (text "sol part 2: ")
  , el [] (text <| String.fromInt <| findPositions2 <| processInput input)
   ]

findPositions : List (Int, Direction) -> Int
findPositions sorted =
  let
    emptyR = Result (0,0) (0,0) Dict.empty Set.empty
  in
    Set.size <| .trail <| moveRope sorted emptyR


findPositions2 : List (Int, Direction) -> Int
findPositions2 sorted =
  let
    rope = Dict.fromList <| List.map (\k-> (k, (0,0))) <| List.range 0 9
    emptyR = Result (0,0) (0,0) rope Set.empty
  in
    Set.size <| .trail <| moveRope2 sorted emptyR

processInput string =
  let
    lines = String.lines string
    dir = Debug.log "directions: " <| makeDirections lines
    sorted = List.sortBy (\(a, _)-> a) dir
  in
    sorted

type alias Direction = {
    steps : Int
    , dir : String
  }

makeDirections lines =
  List.indexedMap processLine lines

processLine : Int-> String -> (Int, Direction)
processLine ind line =
    case String.uncons line of
        Just (d,a) ->
          (ind, Direction (Maybe.withDefault 0 <| String.toInt <| String.trim a) (String.fromChar d))
        Nothing ->
          (ind, Direction 0 "")

type alias Result = {
    head : (Int, Int)
    , tail : (Int, Int)
    , rope : Dict.Dict Int (Int, Int)
    , trail : Set.Set (Int, Int)
  }

moveRope : List (Int, Direction) -> Result -> Result
moveRope dir current =
    case dir of
      (step::rest) ->
          moveRope rest <| Debug.log "processed step " <| processSteps (Debug.log "step " step) current
      []->
        current

moveRope2 : List (Int, Direction) -> Result -> Result
moveRope2 dir current =
    case dir of
      (step::rest) ->
          moveRope2 rest -- <|  Debug.log "2processed step "
          <| processSteps2 step current
      []->
        current

moveHead dir head =
  let
      (x,y) = head
  in
  case dir.dir of
      "U"->
        (x, y + 1)
      "D"->
        (x, y - 1)
      "R"->
        (x + 1, y)
      "L"->
        (x - 1, y)
      _ ->
        head

processSteps : (Int, Direction) -> Result -> Result
processSteps (_, dir) result =
    if dir.steps > 0 then
      let
        newHead : (Int, Int)
        newHead = moveHead dir result.head
        newTail : (Int, Int)
        newTail = moveTail result.tail newHead
        newDir = { dir | steps = dir.steps - 1 }
        newTrail = Set.insert newTail result.trail
        newResult = {result | tail = newTail, head = newHead, trail = newTrail}
      in
      processSteps (0, newDir) newResult
    else
      result

processSteps2 : (Int, Direction) -> Result -> Result
processSteps2 (_, dir) result =
    if dir.steps > 0 then
      let
        newHead : (Int, Int)
        newHead = moveHead dir <| ropeKnot 0 result.rope
        newRopeHead = updateRope result.rope 0 newHead

        newRope = moveNextKnot 0 newRopeHead
        newDir = { dir | steps = dir.steps - 1 }
        newTrail = Set.insert (ropeKnot 9 newRope) result.trail
        newResult = {result | rope = newRope, trail = newTrail}
      in
      processSteps2 (0, newDir) newResult
    else
      result

updateRope rope pos coord =
  Dict.insert pos coord rope

ropeKnot index rope =
  case Dict.get index rope of
    Just head -> head
    Nothing -> (0,0)

moveNextKnot knot rope =
  let
    currentKnot = ropeKnot knot rope
    nextKnotI = (knot + 1)
    nextKnot = ropeKnot nextKnotI rope
    newPosKnot = moveTail nextKnot currentKnot
    newRope = updateRope rope nextKnotI newPosKnot
  in
    if nextKnotI < 9 then
      moveNextKnot nextKnotI newRope
    else
      newRope

-- ........
-- ........
-- ...T.H..
-- ........

moveTail : (Int, Int) -> (Int, Int) -> (Int, Int)
moveTail tail head =
  let
    (xt, yt)= tail
    (xh, yh)= head
    diffX = xh - xt
    diffY = yh - yt
  in
    case (abs diffX,abs diffY) of
      (2, 2) ->
        if diffX == 2 && diffY == 2 then
         ((xt + 1), (yt + 1))
        else
          if diffX == -2 && diffY == 2 then
          ((xt - 1), (yt + 1))
          else
            if diffX == 2 && diffY == -2 then
            ((xt + 1), (yt - 1))
            else
              if diffX == -2 && diffY == -2 then
              ((xt - 1), (yt - 1))
              else
                tail
      (_, 2) ->
        -- Debug.log "abs diff y 2 " <|
        if diffY == 2 then
          -- Debug.log "diff y +2 " <|
            (xh, (yt + 1))
        else
          if diffY == -2 then
            -- Debug.log "diff y -2 " <|
            (xh, (yt - 1))
          else
            -- Debug.log "abs y nope " <|
            tail
      (2,_) ->
        -- Debug.log "abs diff x 2 " <|
        if diffX == 2 then
          -- Debug.log "diff x +2 " <|
            ((xt + 1), yh)
        else
          if diffX == -2 then
            -- Debug.log "diff x -2 " <|
            ((xt - 1), yh)
          else
            tail
      (_,_) ->
        -- Debug.log "no diff " <|
          tail


data1 = """R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2"""


dataR = """R 1
U 2
R 1
L 2
U 2
D 2
L 2
U 2
L 2
R 2
D 2
L 1
D 2
L 2
U 1
L 1
R 2
L 1
R 1
U 2
D 1
U 2
L 1
U 2
D 1
L 1
D 1
L 1
D 2
U 1
D 1
U 1
R 2
U 1
L 2
D 2
U 1
D 2
R 2
L 1
D 1
R 1
U 2
R 2
D 2
L 2
U 1
L 1
U 2
D 1
U 2
R 1
D 1
R 1
L 2
R 2
U 1
D 2
L 2
U 2
D 1
L 1
U 2
D 2
U 1
L 2
R 2
D 1
L 1
D 1
U 2
R 2
D 1
U 1
R 2
D 2
U 2
L 1
R 2
U 1
R 2
L 2
U 2
R 2
L 1
R 1
U 1
R 1
U 1
L 1
D 1
R 1
U 2
D 2
U 2
L 1
D 2
U 1
R 2
D 2
L 1
D 2
L 1
U 1
R 1
L 1
U 1
D 2
U 1
R 2
D 1
U 1
D 1
R 3
L 3
R 2
D 2
U 3
R 3
U 1
D 1
L 3
D 3
L 2
U 3
R 2
U 2
L 3
R 2
D 3
R 1
D 2
R 1
L 1
D 3
R 2
D 1
U 2
L 2
U 1
L 1
D 2
R 2
U 1
R 3
U 2
R 2
D 3
L 2
D 2
R 2
L 3
D 1
U 3
L 2
R 2
L 3
U 2
R 3
U 3
L 2
U 3
R 1
L 3
D 3
U 2
L 1
U 3
R 2
L 1
U 2
R 1
U 1
L 2
R 2
D 1
L 1
U 2
D 1
L 1
D 1
L 1
D 2
U 2
D 3
R 2
D 3
L 3
R 3
U 2
L 2
U 2
L 2
U 2
L 3
U 1
L 2
D 3
L 3
R 3
D 2
U 3
R 1
L 2
R 3
D 1
U 2
R 3
L 3
R 3
U 3
R 2
L 2
R 2
D 2
R 3
D 2
U 3
L 3
D 3
R 3
U 3
R 4
L 3
D 2
L 4
D 2
U 2
L 1
R 3
U 2
R 2
L 4
D 2
R 2
D 2
U 2
R 4
D 1
L 1
U 1
L 3
D 3
U 1
R 1
D 4
U 2
L 3
U 3
D 2
R 3
L 1
R 1
L 1
R 3
L 1
R 1
D 1
R 3
D 3
R 2
U 3
R 2
D 1
L 3
D 2
U 3
D 2
L 4
D 3
R 1
L 4
D 2
R 1
U 3
L 1
R 4
L 3
U 4
R 2
U 1
R 4
L 1
R 1
U 3
D 3
R 3
U 1
L 4
R 4
U 2
R 1
D 2
L 4
R 1
D 2
U 2
L 1
U 4
R 2
D 2
R 3
U 1
L 2
R 3
U 4
R 2
L 1
U 1
R 4
L 2
R 1
L 2
D 3
U 4
L 3
U 4
L 4
U 4
D 3
U 3
L 1
D 1
U 1
D 1
L 4
D 4
L 2
U 4
L 3
D 2
L 1
U 1
D 4
R 1
L 4
U 5
R 5
U 3
D 1
L 4
U 1
L 1
U 2
D 4
R 5
U 5
D 5
L 4
R 4
L 5
U 1
L 1
U 4
L 1
R 5
L 5
R 1
D 2
U 5
R 4
U 3
D 2
U 5
L 3
D 4
U 1
D 2
U 3
R 1
D 2
L 3
U 4
D 1
L 5
D 2
U 3
D 4
R 3
U 1
R 1
D 4
L 5
R 3
L 3
R 5
L 4
R 5
D 4
R 5
U 1
R 2
D 4
R 4
U 2
L 5
R 3
D 5
U 4
D 2
U 5
L 2
R 5
L 3
D 2
R 3
U 5
L 5
D 5
U 2
R 5
U 1
D 2
R 1
L 1
R 3
U 2
D 4
U 1
R 3
U 5
D 5
L 3
U 1
D 1
L 4
R 4
L 4
R 1
U 2
L 3
R 2
U 2
L 1
U 1
L 5
U 3
R 2
U 3
D 1
R 4
D 5
L 1
D 1
U 3
D 2
U 5
L 6
D 2
R 6
D 3
U 4
R 3
U 5
D 4
U 5
D 1
L 3
U 5
R 3
U 1
D 2
R 4
U 3
L 6
D 2
R 4
D 2
L 2
U 6
R 4
U 1
R 4
U 6
L 6
U 4
L 3
D 2
U 4
L 3
R 4
D 4
U 2
L 2
D 1
U 3
L 5
R 3
L 4
D 5
R 1
U 4
L 2
U 4
R 2
L 5
R 6
D 1
U 1
R 6
D 6
U 6
R 6
U 6
R 2
U 6
R 4
D 4
U 5
D 5
R 2
D 1
R 6
D 5
R 2
D 1
L 6
D 5
L 5
D 6
U 5
D 1
U 5
L 1
R 5
L 5
U 5
R 3
U 2
D 6
L 5
U 5
D 5
R 3
D 1
U 6
R 6
D 5
U 5
R 1
D 5
L 3
R 4
U 5
R 5
U 3
D 3
L 5
U 2
L 3
U 2
L 5
R 6
L 1
D 3
U 5
R 5
L 5
U 2
R 4
D 6
L 3
R 2
U 6
D 3
R 6
D 7
L 3
U 2
L 3
D 4
L 2
U 3
D 2
L 4
D 5
R 2
L 3
D 2
R 6
U 6
D 1
R 5
U 1
R 5
U 7
L 6
U 6
L 4
D 7
U 4
R 2
L 7
U 7
L 5
D 4
U 4
R 6
U 4
R 3
U 3
R 4
U 2
R 2
D 4
U 2
R 1
D 2
L 5
D 6
R 6
U 7
L 2
D 4
L 4
D 7
U 7
D 5
L 3
U 5
R 2
D 2
L 6
R 1
D 1
R 5
D 2
R 3
D 5
U 3
L 5
R 7
U 2
R 6
U 5
D 3
R 3
U 4
D 6
U 2
L 4
U 1
R 2
L 6
U 6
R 2
U 6
D 3
L 4
R 1
L 7
D 5
U 3
D 5
R 7
U 1
L 1
U 3
R 5
D 6
R 3
U 7
L 3
R 1
L 7
U 3
D 6
R 4
L 2
D 4
R 6
L 4
D 8
R 3
D 5
L 2
R 2
L 4
U 3
R 1
U 1
D 3
L 1
R 6
U 5
D 6
U 3
D 3
L 7
U 3
D 6
L 2
R 2
D 7
L 7
D 6
L 2
R 3
L 1
R 2
U 3
L 6
D 3
U 5
L 3
U 4
R 8
L 5
U 5
R 3
D 4
R 1
L 7
U 4
R 4
D 8
U 8
D 4
L 8
U 3
L 3
R 4
L 1
U 3
D 4
R 5
L 4
R 3
L 7
R 1
U 1
D 2
L 8
U 5
R 2
L 7
U 1
L 8
R 4
U 3
R 7
U 6
L 3
U 6
L 3
U 3
D 8
L 5
D 7
U 1
L 2
R 4
L 4
R 1
D 8
R 8
L 4
R 5
L 2
R 8
L 7
D 4
L 1
U 1
R 2
U 1
L 2
D 1
L 2
R 8
U 2
L 8
D 5
L 4
R 6
U 1
D 2
L 5
R 7
L 6
R 1
L 2
R 3
L 2
D 2
L 3
R 4
U 7
R 6
D 3
L 4
D 6
R 3
D 6
R 8
U 1
R 6
D 1
U 7
D 9
R 2
D 1
U 6
R 3
D 5
U 5
L 6
U 3
D 7
L 2
D 6
L 5
U 4
L 2
D 5
L 3
D 4
U 7
L 4
D 9
R 3
L 5
R 5
L 5
U 1
R 8
L 6
U 2
D 2
R 3
L 5
D 7
L 4
U 4
D 9
L 8
R 1
L 5
R 8
D 9
R 6
L 6
U 6
R 3
U 5
R 8
U 4
L 8
U 3
R 2
L 7
R 9
U 6
L 2
U 2
D 9
U 6
R 5
D 8
R 1
D 1
L 4
U 2
D 9
L 9
D 5
U 9
R 2
D 6
R 9
L 1
R 2
U 1
L 7
R 4
U 4
L 2
U 2
L 5
D 5
L 6
D 2
R 8
L 4
D 9
R 8
L 2
D 4
R 2
L 6
R 10
U 7
L 4
U 9
R 8
D 9
U 10
R 2
L 4
U 7
D 8
R 8
D 5
L 5
D 5
L 2
D 1
U 5
D 6
R 7
L 10
U 9
R 3
D 9
U 2
L 8
U 3
D 2
U 5
D 5
R 2
U 5
R 4
L 8
U 10
L 7
D 4
L 1
U 7
L 10
R 1
L 6
D 10
R 2
U 3
D 1
R 2
D 2
U 6
L 1
D 2
U 5
D 6
L 4
D 3
R 4
L 3
D 10
U 2
R 8
U 5
L 5
R 8
D 6
L 8
U 8
R 9
U 3
L 9
U 2
R 5
U 1
R 1
D 2
L 3
R 9
D 10
L 2
R 9
D 4
L 7
U 7
R 2
D 10
U 5
R 9
L 6
D 4
U 10
R 5
D 2
R 7
U 10
L 3
D 9
L 3
R 6
L 8
U 5
R 10
L 9
U 4
R 8
D 2
R 4
U 10
L 1
U 9
L 9
R 3
D 4
L 7
R 6
D 8
L 5
R 9
L 8
R 7
D 10
L 3
D 3
U 6
L 5
R 5
L 9
R 2
D 2
U 5
R 8
D 11
U 10
L 3
D 6
U 11
R 11
U 4
D 7
L 8
R 10
U 10
L 5
R 7
L 4
U 2
L 5
U 3
D 6
U 11
D 6
L 1
D 10
L 6
D 9
R 3
D 5
R 6
U 5
L 7
R 9
D 5
R 11
D 10
U 1
L 2
R 8
L 6
D 5
U 10
L 10
R 5
U 2
D 3
L 2
R 6
D 6
L 1
U 5
R 8
D 7
R 10
L 4
D 2
R 10
U 8
L 10
R 7
U 8
D 7
U 3
R 8
D 10
U 11
D 11
L 4
U 5
R 10
U 2
L 11
U 10
R 10
U 1
D 6
U 1
L 3
R 11
L 2
U 3
D 1
U 11
L 2
D 4
R 9
L 2
D 7
U 1
R 2
L 8
R 6
D 4
R 1
L 6
D 4
U 8
D 8
R 5
U 11
R 6
U 11
R 11
U 7
R 9
D 4
R 1
D 9
R 3
L 9
U 4
R 6
D 5
L 3
R 11
U 11
L 6
R 11
L 7
D 1
R 10
D 2
U 8
R 4
L 6
R 4
U 4
D 12
R 6
L 12
R 9
L 8
R 2
L 12
U 2
L 9
D 2
U 8
D 9
R 12
U 6
L 2
U 11
R 5
D 7
U 11
D 1
R 6
U 3
D 9
L 2
R 1
D 6
R 12
D 5
L 11
R 11
L 3
R 11
L 1
D 11
U 10
L 2
R 2
U 8
D 3
U 3
L 4
D 8
L 2
D 9
R 10
D 9
U 11
L 3
D 4
L 5
U 4
R 5
D 12
U 2
L 11
R 1
L 6
D 11
R 9
U 10
D 2
L 12
D 7
U 9
D 12
U 5
L 4
D 7
L 11
D 7
L 2
D 10
R 7
D 1
R 9
U 4
D 12
U 4
R 3
D 1
L 6
D 3
R 9
L 12
D 10
U 2
L 12
D 6
L 2
D 5
L 1
U 4
L 5
U 5
R 7
U 9
R 4
D 9
L 13
U 2
L 8
U 1
R 10
L 5
U 2
L 7
U 6
L 10
U 11
R 2
U 6
R 13
D 4
U 3
R 11
D 2
U 7
D 4
R 13
L 7
R 11
L 8
U 4
R 13
D 9
U 11
L 4
D 10
U 8
D 8
R 6
U 7
L 8
R 2
U 7
L 12
R 2
D 8
L 7
U 6
D 6
L 9
R 8
L 11
U 6
D 3
U 2
L 2
R 1
U 7
L 6
U 5
D 6
U 11
D 11
U 8
R 12
L 7
R 3
D 8
L 11
U 6
D 10
L 2
D 6
L 5
D 10
U 12
D 6
R 10
D 4
R 13
U 2
D 3
L 9
U 12
L 10
D 11
R 7
D 5
L 5
D 2
L 11
U 6
L 1
U 11
D 5
R 8
U 1
D 11
U 8
L 6
D 6
U 12
L 7
R 14
D 8
R 11
D 12
R 7
U 6
D 12
R 6
D 6
L 6
U 14
D 2
U 11
D 11
L 14
U 8
L 12
D 12
U 9
R 7
D 2
U 10
D 8
R 8
D 8
L 10
D 11
L 4
D 11
U 4
D 6
U 5
D 8
R 14
D 5
R 1
L 8
R 8
L 2
R 11
D 9
L 3
R 7
L 5
U 14
L 1
R 8
L 12
R 5
D 11
U 6
R 1
U 10
R 11
U 11
L 9
R 9
L 8
R 1
D 3
R 11
U 12
D 14
L 1
U 4
D 2
L 11
R 4
U 14
R 1
L 3
R 13
L 13
D 8
L 14
U 12
R 13
D 9
R 8
D 12
U 14
D 2
L 10
R 1
D 12
R 13
L 1
U 5
L 12
R 6
U 14
R 6
L 13
U 11
R 2
D 6
R 5
U 12
R 1
D 11
L 4
U 6
R 4
U 5
L 12
D 3
R 4
D 1
L 15
R 15
U 13
R 10
U 14
R 15
D 4
R 8
U 13
L 4
R 5
D 7
R 5
L 2
D 7
L 4
R 7
D 6
R 3
D 12
U 9
L 14
R 7
L 12
R 7
D 15
R 4
L 1
D 11
R 10
D 8
L 9
D 10
L 4
R 10
U 8
D 2
L 8
D 7
U 11
L 9
U 15
L 7
R 6
L 14
D 2
L 15
D 6
U 2
R 9
D 6
U 7
R 7
U 15
D 13
L 6
D 9
R 12
U 2
R 8
L 14
R 11
D 7
U 1
R 11
D 12
L 15
D 9
U 7
L 3
U 12
L 3
D 9
R 12
U 2
D 8
L 1
U 15
R 15
U 15
L 10
U 2
D 4
R 8
U 4
D 8
L 7
R 11
U 15
D 9
R 10
D 8
R 6
D 6
U 3
L 13
D 10
L 11
D 7
U 5
D 8
U 7
R 10
U 1
L 5
D 13
L 13
U 11
R 14
D 7
L 4
R 8
D 12
U 14
D 7
L 5
D 14
U 2
L 16
R 9
U 13
R 14
L 16
U 16
L 9
U 13
R 5
U 15
R 6
D 7
R 15
L 15
U 3
L 5
R 3
U 1
R 14
U 5
L 6
U 13
D 6
U 1
R 10
D 11
L 16
U 10
R 16
D 15
L 7
R 15
U 15
R 5
D 9
U 6
D 7
L 15
R 9
L 10
U 13
R 7
D 15
U 14
D 12
L 2
U 11
D 14
R 4
L 1
U 1
R 13
D 12
R 14
L 4
R 11
U 14
L 2
D 6
R 15
U 12
R 15
U 9
D 11
R 4
L 15
U 12
D 14
L 11
U 10
L 15
U 9
D 7
R 12
L 14
U 15
L 6
U 2
R 11
L 12
R 5
U 8
D 7
U 6
R 2
U 6
R 16
L 2
R 11
L 7
D 16
R 7
L 4
R 13
U 12
D 1
R 2
D 4
L 10
D 12
R 3
D 14
R 17
L 4
D 3
L 2
R 13
L 15
R 13
U 9
D 7
L 13
U 14
D 6
L 6
D 6
R 3
U 6
R 8
U 14
R 13
L 15
D 2
U 3
D 10
R 16
L 10
D 8
R 10
D 9
U 8
R 1
L 4
U 14
R 7
L 16
R 8
D 4
R 9
U 12
D 8
U 7
R 9
L 16
D 15
U 11
D 15
U 2
R 9
U 10
L 14
D 14
R 4
U 12
L 1
U 8
R 3
U 13
D 17
U 17
D 6
U 17
L 11
D 16
L 10
R 10
L 6
U 16
L 7
D 8
R 16
U 7
D 9
L 2
U 5
L 4
U 10
D 1
L 15
U 5
L 14
R 7
U 15
L 14
U 14
L 7
D 3
L 4
U 1
R 11
D 12
R 9
U 16
L 3
D 3
L 4
U 5
D 16
U 1
D 15
R 1
L 16
U 7
L 2
U 12
L 13
U 16
D 9
R 4
D 15
R 7
U 14
L 2
U 9
R 2
L 13
U 16
R 4
D 17
U 10
D 10
U 10
D 18
R 10
U 7
L 6
U 9
D 9
L 5
U 9
R 9
L 15
U 9
D 10
U 10
R 10
D 2
U 15
D 3
R 18
U 13
R 1
D 17
L 16
D 10
L 1
R 7
U 1
D 12
L 3
U 12
L 6
R 18
D 13
R 18
U 17
D 7
L 11
D 14
L 2
U 18
D 8
L 6
D 16
U 4
R 14
U 3
D 11
R 12
U 12
R 5
D 16
R 17
U 3
D 15
U 11
L 17
U 15
L 17
U 12
D 15
L 14
D 14
U 13
L 16
R 5
U 16
D 6
U 13
L 13
D 4
R 6
D 6
L 18
U 18
L 15
R 6
L 6
R 6
U 13
L 8
U 13
R 13
U 13
L 15
R 6
U 8
R 10
U 11
L 18
U 3
L 2
D 8
L 5
U 11
L 5
U 8
R 13
U 12
R 9
U 18
D 11
L 12
D 13
R 18
U 2
L 9
D 14
R 2
L 17
U 3
L 9
U 10
R 9
L 13
R 15
D 14
U 2
R 4
L 5
U 12
L 14
U 15
D 13
R 10
L 5
U 15
D 10
R 7
L 9
R 15
L 6
U 9
L 15
U 10
R 14
L 17
R 16
U 12
D 15
L 5
U 17
R 12
D 17
U 8
R 10
L 9
R 6
U 1
D 3
L 3
D 7
U 6
L 9
U 19
R 10
U 14
R 15
U 4
D 9
R 17
L 6
U 9
R 2
U 9
R 16
L 13
U 11
R 15
D 10
L 5
U 1
R 4
U 17
L 1
U 10
L 3
U 9
L 17
U 8
L 2
U 10
D 12
L 16
R 18
L 6
D 11
R 6
U 7
R 3
L 15
R 7
U 14
R 2
U 7
R 17
L 7
R 1
D 11
R 10
L 17
R 5
U 10
L 15
U 7
R 7
L 17
D 19
U 17
R 3
D 6
L 8
R 8
D 3
U 7
L 9"""

data2 = """R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20"""
