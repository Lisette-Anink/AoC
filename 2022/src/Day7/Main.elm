module Day7.Main exposing (..)

import Browser
import Html exposing (Html, button, div, text, p)
import Html.Attributes exposing (style, class)
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
  div [][
    -- div [class "sol", style "display" "flex"]
    --   <| part2 data1
    -- , div [class "sol", style "display" "flex"]
    --   <| part2 dataR
    -- , div [class "sol", style "display" "flex"]
      -- <| part1 data1
    -- , div [class "sol", style "display" "flex"]
    --   <| part1 data2
    -- , div [class "sol", style "display" "flex"]
    --   <| part1 data3
    -- , div [class "sol", style "display" "flex"]
    --   <| part1 dataR
  ]

-- part1 input =
--   [p [] [text "sol part 1: "]
--   , p [] [text <| String.fromInt <| processInput input]]


-- processInput : String -> Int
-- processInput input =
--   let
--     lines = String.lines input
--     fSizes = Dict.empty
--     d = findFolderSize fSizes lines

--   in
--     d


-- findFolderSize dict lines =
--   let
--     isDir = String.startsWith "dir"
--     isCommand = String.startsWith "$"

--   in
--     case lines of
--       (first:: rest) ->
--         -- let
--           -- newDict = updateDict first dict
--         -- in
--           findFolderSize newDict rest
--       ([]) ->
--         dict


-- updateDict string dict =
--   -- let

--   -- in
--     Dict.update f addSize dict


sumOfSmall dict =
  0

data1 = """$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k"""


dataR = """"""

