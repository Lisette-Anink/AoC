module Day5.Main exposing (..)

import Browser
import Element exposing (Element, el, text, row,  column, alignRight, fill, width, rgb255, spacing, centerY, padding)
import Element.Background as Background
import Element.Border as Border
import Element.Font as Font
import Element.Events as Events
import Dict
import List exposing (drop, take)
import Set
import Element exposing (alignBottom)
import Html exposing (button)
import Regex
import Task
import Process
import Day2.Main exposing (subscriptions)


main =
  Browser.element { init = init, update = update, view = view, subscriptions = subscriptions }

type alias Model =
  {cols : Dict.Dict Int (List Char)
  , dataMove : String
  , dataStack : String
  , moves : Dict.Dict Int Move
  , currentMove : Int
  }

subscriptions : Model -> Sub Msg
subscriptions model =
    Sub.none


init : () -> (Model, Cmd Msg)
init _ =
  (Model (parseCols dataS) dataM dataS Dict.empty 0
  , Cmd.none )

type Msg = Start
  | Step

update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case Debug.log "msg " msg of
    Start ->
      (parseInstructions model, processMove)
    Step ->
      if model.currentMove < (List.length <| Dict.values model.moves) then
        let
          newModel = moveCrate model
        in
          ({newModel | currentMove = model.currentMove + 1}, Cmd.none)
      else
        (model, Cmd.none)


view model =
    let
       vals = Dict.values model.cols
    in
    Element.layout []
        <| column [spacing 10]
           [ topCrates vals
           , el [] (text <| "currentMove: " ++ String.fromInt model.currentMove)
          , row []
            <| List.map myRowOfStuff vals
          , el [Events.onClick Start] (text "start")
          , movesView model.moves
           ]

myRowOfStuff : List Char -> Element msg
myRowOfStuff stack =
    column [ width fill, alignBottom, spacing 3 ]
        <| showStack (List.reverse stack)

showStack list =
  case list of
      (top::rest) ->
         myElementT (String.fromChar top) :: showStack rest
      [] ->
        []

topV : List Char -> Char
topV list =
  case List.reverse list of
    (t::r) -> t
    [] -> ' '


topCrates : List (List Char) -> Element.Element msg
topCrates listList =
  let
    topCrate = List.map topV listList
  in
  el [] (text <| String.fromList topCrate)

myElementT : String -> Element msg
myElementT string =
    el
        [ Background.color (rgb255 240 0 245)
        , Font.color (rgb255 255 255 255)
        , Border.rounded 3
        , padding 3

        ]
        (text string)

myElement : String -> Element msg
myElement string =
    el
        [
        Border.rounded 3
        , padding 3
        ]
        (text string)

movesView : Dict.Dict Int Move -> Element msg
movesView movesD =
  let
      moveVs = Dict.values movesD
  in
    column []
      <| el [] (text "moves : ")
      :: List.map (\m-> el [] <| text ("move "++ String.fromInt m.amount ++"from"++ String.fromInt m.from ++ "to"++ String.fromInt m.to)) moveVs



processMove =
    delay 2000 <| Step


moveCrate : Model -> Model
moveCrate model =
  let
    cols = model.cols
    move = Dict.get model.currentMove model.moves
  in
    case move of
      Just mov ->
        let
          newCols = moveC cols mov
        in
          {model | cols = newCols}
      Nothing -> model

moveC : Dict.Dict Int (List Char) -> Move -> Dict.Dict Int (List Char)
moveC cols move =
  let
    from = Dict.get move.from cols
    to = Dict.get move.to cols
  in
    case (from, to) of
        (Just listF, Just listT) ->
          let
            (newF, newT) = moveTop 0 move.amount listF listT
            newDT = Dict.insert move.to newT cols
          in
            Dict.insert move.from newF newDT

        (_, _) -> cols

moveTop : Int -> Int -> List Char -> List Char -> (List Char, List Char)
moveTop count total colF colT =
  if count < total then
    case List.reverse colF of
      (t::r) ->
        let
          newT = (List.append colT [t])
          newF = List.reverse r
        in
          moveTop (count + 1) total newF newT
      _ -> (colF, colT)
  else
    (colF, colT)


delay : Float -> msg -> Cmd msg
delay time msg =
  Process.sleep time
  |> Task.perform (\_ -> msg)

parseInstructions : Model -> Model
parseInstructions model =
  let
    moveLines = String.lines model.dataMove
    p = Dict.fromList <| List.indexedMap Tuple.pair <| List.map parseMove moveLines
  in
    { model | moves = p }

type alias Move = {amount: Int, from: Int, to: Int}

parseMove line =
  -- move 3 from 1 to 3
  let
      regex = Maybe.withDefault Regex.never <|
            Regex.fromString "\\s(\\d*)"
      matches = Regex.find regex line
      vals = List.map .match matches
      intVals = List.filterMap (\v-> String.toInt <| String.trim v) vals
  in
    case intVals of
      (a::f::t::r) ->
          let
              string = Debug.log "parse move: "  (a , f ,  t)
          in

          Move a f t
      _ -> Move 0 0 0


parseCols string =
  let
    lines = String.lines string
    linesToChar = transpose <| List.map String.toList lines
    dict = emptyDict
  in
    createDict linesToChar

emptyDict : Dict.Dict Int (List Char)
emptyDict =
    Dict.singleton 0 []

createDict : List (List Char) -> Dict.Dict Int (List Char)
createDict listList =
    case listList of
      (first::rest) ->
        addToDict (Debug.log "row: " (List.reverse first)) (createDict rest)
      [] ->
        emptyDict

addToDict : List Char -> Dict.Dict Int (List Char) ->  Dict.Dict Int (List Char)
addToDict list dict =
  case list of
    (k::v) ->
      let
          i = String.toInt <| String.fromChar k
      in
      case i of
        Just int ->
          Dict.insert int (removeEmpty v) dict
        Nothing ->
          dict
    [] -> dict

removeEmpty list =
  List.filter (
    \n-> not <| String.isEmpty
        <| String.trim
        <| String.fromChar n
    ) list

transpose : List (List a) -> List (List a)
transpose ll =
  let heads = List.map (List.take 1) ll |> List.concat
      tails = List.map (List.drop 1) ll
  in
      if List.isEmpty heads then
             []
         else
             heads::(transpose tails)


-- part 2 stuff


dataS = """
    [D]
[N] [C]
[Z] [M] [P]
 1   2   3"""

dataM = """move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2"""

dataR = """"""

