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
          <| part1 data1
            ++ part1 data1

part1 input =
  [el [] (text "sol part 1: ")
  , el [] (text <| String.fromInt <| findVisibleTrees <| processInput input)
   ] ++  (showVisibleTrees <| processInput input)


processInput string =
  let
    lines = String.lines string
  in
    forest

processLine : (Int, String) -> (Int, List (Int, Int))
processLine (ind, line) =
  let
    indexedLine =
      List.indexedMap Tuple.pair
        <| List.filterMap String.toInt
        <| String.split "" line
  in
    (ind, indexedLine)






data1 = """"""


dataR = """"""

