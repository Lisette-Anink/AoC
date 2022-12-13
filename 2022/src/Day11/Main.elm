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
            <| highestInspectCount
            <| processInput input)
   ]

processInput string =
  let
    lines = String.split "\n\n" string
    monkeys = Debug.log "monkeys: " <| makeMonkeys Dict.empty lines
  in
    Debug.log "after play " <| playKeepAway monkeys 0

highestInspectCount : Dict.Dict String Monkey -> Int
highestInspectCount monkeys =
  let
    vals = Debug.log "inspectCounts " <| List.reverse <| List.sort <| List.map .inspectCount <| Dict.values monkeys
  in
    case vals of
      (a::b::_) ->
        a * b
      (_) ->
        0

playKeepAway monkeys r =
  let
      newRound = r + 1
      keys = Dict.keys monkeys

  in
    if newRound < 21 then
      playKeepAway ( Debug.log "after round: " <| playRound monkeys keys) newRound
    else
      monkeys

playRound : Dict.Dict String Monkey -> List String -> Dict.Dict String Monkey
playRound monkeys keys =
    case keys of
      (first::rest) ->
        let
          monkey = Dict.get first monkeys
        in
          case monkey of
              Just monk ->
                  playRound (processMonkey first monk monkeys) rest
              Nothing ->
                Debug.log "Something is wrong! " monkeys
      (_)->
        monkeys

processMonkey : String -> Monkey -> Dict.Dict String Monkey -> Dict.Dict String Monkey
processMonkey current monkey monkeys=
  let
    items = monkey.items
  in
  case items of
    (f::r) ->
      let
          newItems = {monkey | items = r, inspectCount = monkey.inspectCount + 1}
          uMonkeys = Dict.insert current newItems monkeys

          updatedMonkeys = updateMonkey f newItems uMonkeys
      in
      processMonkey current newItems updatedMonkeys
    [] ->
      monkeys

updateMonkey : Int -> Monkey -> Dict.Dict String Monkey -> Dict.Dict String Monkey
updateMonkey item monkey monkeys =
  let
    (newKey, newLevel) = updateWorry item monkey
    other = Dict.get newKey monkeys
  in
    case other of
      Just a ->
        let
          valM = {a | items =  a.items ++ [newLevel] }
        in
          Dict.insert newKey valM monkeys
      Nothing ->
        monkeys

updateWorry : Int -> Monkey -> (String, Int)
updateWorry item monkey =
  let
    -- m = Debug.log "Monkey nr " <| monkey.nr
    level = case monkey.operation of
        Add x ->
          (item + x) --// 3
        Multiply x ->
          (item * x) --// 3
        MultiplySelf ->
          (item * item) --// 3
    remainder = remainderBy monkey.test level
  in
    if   remainder == 0 then
      -- Debug.log "Throw to ifTrue " <|
      (monkey.ifTrue, item + remainder)
    else
      -- Debug.log "Throw to ifFalse " <|
      (monkey.ifFalse, item + remainder)


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
    nr : String
    , items : List Int
    , inspectCount : Int
    , operation : Operation
    , test : Int
    , ifTrue : String
    , ifFalse : String
  }


matchOperation line =
  let
      isAdd = String.contains "+" line
      isMultiply = String.contains "*" line
      nr = Maybe.withDefault 0 <| String.toInt <| String.filter Char.isDigit line
  in
    if isAdd then
      Debug.log "IS +" Add nr
    else
      if isMultiply then
        if nr == 0 then
          MultiplySelf
        else
          Multiply nr
      else
        Debug.log "NOTHING WHAT!" MultiplySelf


makeMonkey dict string  =
  let
      lines =  List.indexedMap Tuple.pair <| String.lines string
      monkeyNr = String.filter Char.isDigit <| getLine lines 0
      items = List.filterMap (\v-> String.toInt (String.filter Char.isDigit v))
                      <| String.split "," <| getLine lines 1
      op = matchOperation <| getLine lines 2
      test = Maybe.withDefault 0 <| String.toInt <| String.filter Char.isDigit <| getLine lines 3
      ifTrue = String.filter Char.isDigit <| getLine lines 4
      ifFalse = String.filter Char.isDigit <| getLine lines 5
      monkey = Monkey monkeyNr items 0 op test ifTrue ifFalse
  in
     Debug.log "insert monkey: " <| Dict.insert monkeyNr monkey dict

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

dataR = """Monkey 0:
  Starting items: 73, 77
  Operation: new = old * 5
  Test: divisible by 11
    If true: throw to monkey 6
    If false: throw to monkey 5

Monkey 1:
  Starting items: 57, 88, 80
  Operation: new = old + 5
  Test: divisible by 19
    If true: throw to monkey 6
    If false: throw to monkey 0

Monkey 2:
  Starting items: 61, 81, 84, 69, 77, 88
  Operation: new = old * 19
  Test: divisible by 5
    If true: throw to monkey 3
    If false: throw to monkey 1

Monkey 3:
  Starting items: 78, 89, 71, 60, 81, 84, 87, 75
  Operation: new = old + 7
  Test: divisible by 3
    If true: throw to monkey 1
    If false: throw to monkey 0

Monkey 4:
  Starting items: 60, 76, 90, 63, 86, 87, 89
  Operation: new = old + 2
  Test: divisible by 13
    If true: throw to monkey 2
    If false: throw to monkey 7

Monkey 5:
  Starting items: 88
  Operation: new = old + 1
  Test: divisible by 17
    If true: throw to monkey 4
    If false: throw to monkey 7

Monkey 6:
  Starting items: 84, 98, 78, 85
  Operation: new = old * old
  Test: divisible by 7
    If true: throw to monkey 5
    If false: throw to monkey 4

Monkey 7:
  Starting items: 98, 89, 78, 73, 71
  Operation: new = old + 4
  Test: divisible by 2
    If true: throw to monkey 3
    If false: throw to monkey 2"""
