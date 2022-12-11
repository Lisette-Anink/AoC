module Day10.Main exposing (..)

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
          part1 dataR
            ++ part1 data2
          --  part2 data1
          -- ++ part2 data2
          -- ++ part2 dataR

part1 input =
  let
    result =  processInput input
    t = (String.fromInt result.cycle) ++ "with x " ++ (String.fromInt result.x)
    listProduct = List.map2 (*) (Dict.values result.strength) (Dict.keys result.strength)
    showProduct : Int -> Int -> Element msg
    showProduct k v = el [] (text  (" " ++ String.fromInt  (k * v)))

    final = (String.fromInt <| List.sum listProduct)
  in

  [el [] (text "sol part 1: ")
  , el [] (text t )
  , el [] (text final )] ++
   List.map (\v-> el [] (text <| " "++String.fromInt v))  (Dict.keys result.strength) ++
   List.map (\v-> el [] (text <| " "++String.fromInt v))  (Dict.values result.strength) ++
   List.map2 showProduct (Dict.values result.strength) (Dict.keys result.strength)




processInput string =
  let
    lines = String.lines string
    result = Debug.log "directions: " <| makeDirections lines
  in
    result

makeDirections lines =
  processLine (Result 0 1 20 Dict.empty) lines

type alias Result = {
    cycle : Int
    , x : Int
    , strengthCycle : Int
    , strength : Dict.Dict Int Int
  }

processLine : Result -> List String -> Result
processLine result lines =
  case lines of
    (first::rest) ->
        let
          w = String.words first
        in
        case reverse w of
            (f::_) ->
                case String.toInt f of
                  Just n ->
                      let
                        r = checkStrength result 2
                        newR = {r | cycle = result.cycle + 2, x =  result.x + n}
                      in
                      processLine newR rest
                  Nothing ->
                      let
                        r = checkStrength result 1
                        newR = {result | cycle = result.cycle + 1}
                      in
                      processLine newR rest
            [] ->
              Debug.log "This should NOT HAPPEN!" result
    [] ->
      result

checkStrength : Result -> Int -> Result
checkStrength result add =
    if result.cycle + add >= result.strengthCycle then
      let
        newStrength = Debug.log "Yes! " Dict.insert result.strengthCycle result.x  result.strength
      in
        { result | strengthCycle = result.strengthCycle +40, strength = newStrength }
    else
      -- Debug.log "noop"
      result


data1 = """noop
addx 3
addx -5"""

data2 = """addx 15
addx -11
addx 6
addx -3
addx 5
addx -1
addx -8
addx 13
addx 4
noop
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx 5
addx -1
addx -35
addx 1
addx 24
addx -19
addx 1
addx 16
addx -11
noop
noop
addx 21
addx -15
noop
noop
addx -3
addx 9
addx 1
addx -3
addx 8
addx 1
addx 5
noop
noop
noop
noop
noop
addx -36
noop
addx 1
addx 7
noop
noop
noop
addx 2
addx 6
noop
noop
noop
noop
noop
addx 1
noop
noop
addx 7
addx 1
noop
addx -13
addx 13
addx 7
noop
addx 1
addx -33
noop
noop
noop
addx 2
noop
noop
noop
addx 8
noop
addx -1
addx 2
addx 1
noop
addx 17
addx -9
addx 1
addx 1
addx -3
addx 11
noop
noop
addx 1
noop
addx 1
noop
noop
addx -13
addx -19
addx 1
addx 3
addx 26
addx -30
addx 12
addx -1
addx 3
addx 1
noop
noop
noop
addx -9
addx 18
addx 1
addx 2
noop
noop
addx 9
noop
noop
noop
addx -1
addx 2
addx -37
addx 1
addx 3
noop
addx 15
addx -21
addx 22
addx -6
addx 1
noop
addx 2
addx 1
noop
addx -10
noop
noop
addx 20
addx 1
addx 2
addx 2
addx -6
addx -11
noop
noop
noop"""

dataR = """addx 2
addx 4
noop
noop
addx 17
noop
addx -11
addx -1
addx 4
noop
noop
addx 6
noop
noop
addx -14
addx 19
noop
addx 4
noop
noop
addx 1
addx 4
addx -20
addx 21
addx -38
noop
addx 7
noop
addx 3
noop
addx 22
noop
addx -17
addx 2
addx 3
noop
addx 2
addx 3
noop
addx 2
addx -8
addx 9
addx 2
noop
noop
addx 7
addx 2
addx -27
addx -10
noop
addx 37
addx -34
addx 30
addx -29
addx 9
noop
addx 2
noop
noop
noop
addx 5
addx -4
addx 9
addx -2
addx 7
noop
noop
addx 1
addx 4
addx -1
noop
addx -19
addx -17
noop
addx 1
addx 4
addx 3
addx 11
addx 17
addx -23
addx 2
noop
addx 3
addx 2
addx 3
addx 4
addx -22
noop
addx 27
addx -32
addx 14
addx 21
addx 2
noop
addx -37
noop
addx 31
addx -26
addx 5
addx 2
addx 3
addx -2
addx 2
addx 5
addx 2
addx 3
noop
addx 2
addx 9
addx -8
addx 2
addx 11
addx -4
addx 2
addx -15
addx -22
addx 1
addx 5
noop
noop
noop
noop
noop
addx 4
addx 19
addx -15
addx 1
noop
noop
addx 6
noop
noop
addx 5
addx -1
addx 5
addx -14
addx -13
addx 30
noop
addx 3
noop
noop"""
