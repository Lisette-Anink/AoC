module Day11.Main exposing (..)

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
import Parser as P exposing ((|.), (|=))
import Regex

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
          part1 data1
            -- ++ part1 dataR
          --  part2 data1
          -- ++ part2 data2
          -- ++ part2 dataR

part1 input =
  [el [] (text "sol part 1: ")
  , el [] (text <| String.fromInt 
            -- <| findPositions 
            <| processInput input)
   ]

processInput string =
  let
    lines = String.split "/n/n" string
    monkeys = Debug.log "monkeys: " <| makeMonkeys Dict.empty lines
    
  in
    0

type Operation 
  = Multiply Int
  | Add Int
  | MultiplySelf


makeMonkeys dict lines =
    case lines of 
      (first::rest)->
        let
          newDict = makeMonkey dict first            
        in
          makeMonkeys newDict rest
      []-> dict

type alias Monkey = {
    items : List Int
    , operation : Operation
    , test : Int
    , ifTrue : Int 
    , ifFalse : Int
  }
  -- Monkey 0:
  -- Starting items: 79, 98
  -- Operation: new = old * 19
  -- Test: divisible by 23
  --   If true: throw to monkey 2
  --   If false: throw to monkey 3



matchOperation line =
  let
      r =  Maybe.withDefault Regex.never <|
              Regex.fromString ".*(\\+|\\*).*(\\d+|old)"
      submatches = List.map (\s-> Debug.log "subm" <| .submatches s) <| Regex.find r <| Debug.log "l: " line
  in
    case submatches of
        (sign::_) ->
          let
              l = List.filterMap identity sign
              lm = List.filterMap identity sign
          in
            case l of
                (a::_) ->
                  if a == "+" then
                    Debug.log "IS +" Add 0
                  else
                    Debug.log "IS *" MultiplySelf
                (_)->
                  Debug.log "NoSign !!" MultiplySelf

  
        (_) ->
          Debug.log "THIS SHOULD NOT HAPPEN!" MultiplySelf
  

makeMonkey dict string  =
  let
      lines =  List.indexedMap Tuple.pair <| String.lines string
      monkeyNr = String.filter Char.isDigit <| getLine lines 0
      items = List.filterMap (\v-> String.toInt (String.filter Char.isDigit v)) 
                      <| String.split "," <| getLine lines 1
      op = matchOperation <| getLine lines 2
      test = Maybe.withDefault 0 <| String.toInt <| String.filter Char.isDigit <| getLine lines 3
      ifTrue = Maybe.withDefault 0 <| String.toInt <| String.filter Char.isDigit <| getLine lines 4
      ifFalse = Maybe.withDefault 0 <| String.toInt <| String.filter Char.isDigit <| getLine lines 5
      monkey = Monkey items op test ifTrue ifFalse
  in
    Dict.insert monkeyNr monkey dict

getLine : List (Int, String) -> Int -> String
getLine lines index =
  let 
    d = Dict.fromList lines
  in 
    Maybe.withDefault "" <| Dict.get index d


data1 = """Monkey 0:
  Starting items: 79, 98
  Operation: new = old * 19
  Test: divisible by 23
    If true: throw to monkey 2
    If false: throw to monkey 3

Monkey 1:
  Starting items: 54, 65, 75, 74
  Operation: new = old + 6
  Test: divisible by 19
    If true: throw to monkey 2
    If false: throw to monkey 0

Monkey 2:
  Starting items: 79, 60, 97
  Operation: new = old * old
  Test: divisible by 13
    If true: throw to monkey 1
    If false: throw to monkey 3

Monkey 3:
  Starting items: 74
  Operation: new = old + 3
  Test: divisible by 17
    If true: throw to monkey 0
    If false: throw to monkey 1"""


dataR = """"""
