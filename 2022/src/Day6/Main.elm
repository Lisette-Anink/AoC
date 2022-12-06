module Day6.Main exposing (..)

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


type Msg = Start | Step

update msg model =
  case msg of
    Start ->
      model
    Step ->
      model


view model =
  div [][
    div [class "sol", style "display" "flex"]
      <| part2 data1
    , div [class "sol", style "display" "flex"]
      <| part2 dataR
    , div [class "sol", style "display" "flex"]
      <| part1 data1
    , div [class "sol", style "display" "flex"]
      <| part1 data2
    , div [class "sol", style "display" "flex"]
      <| part1 data3
    , div [class "sol", style "display" "flex"]
      <| part1 dataR
  ]

part1 input =
  [p [] [text "sol part 1: "]
  , p [] [text <| String.fromInt <| returnIndex input 4 4]]

part2 input =
  [p [] [text "sol part 2: "]
  , p [] [text <| String.fromInt <| returnIndex input 14 14]]


returnIndex : String -> Int -> Int -> Int
returnIndex str i n =
  if fourUniq <| String.slice (i - n) i str then
    i
    else
    returnIndex str (i + 1) n


fourUniq : String -> Bool
fourUniq str =
  let
    m = List.all (\c-> (List.length <| String.indexes c str) <= 1) <| String.split "" str
  in
    m


data1 = """mjqjpqmgbljsphdztnvjfqwrcgsmlb"""
data2 = """bvwbjplbgvbhsrlpgdmjqwftvncz"""
data3 = """zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"""


dataR = """jghhttcmttdwwfjjjpqjqllwvwffswwqmwmddvndvdrrcwrccfncfcgcjjpjrjzjggczgglhllbzlztlzzswwbqqrvrvqrqcrclrccrbrpbbzhbbzvbbwrwswqwqqgmmgqmqmdqdhdmmfnndpdvvdvdbbmgmqgqgqmqrrsbsrrzcrzrcrvrddpqdpdwwsvshvshvshszhshddjllbjljjwhjwhwppbssdccvzzmvvwrrrhmmbdbggqhgqglglgrrdbdjddjvddzpznzmnnpgnppnznzvzbzvvdtdwdttnzzdpzzqjzjzllhglghgttcbtcbtbjbssqcqtcqtqvttvpvqvggbsszjsjqjllfgfqqhllvsvmmjhhbdbzzrwzzggdfgdgcddgmmjnjhhhgvgwghgfffclffdmmmlclzclzzqgqjqzqwzqzbbbljlplnlrlqlqlhlvlmvvlwltwtqwwznndpdttjjrllpptnngjnnzppjhpjhpjpjmjpmjpptzppjpsprpbpmmshmsmqqnhnjhnjnccdnnqhhgfgmgnmnvvtwvttcbtbnbcnnhjjnpphjpjppshsppmrmrbrlrzlzjlzzvssvfffsjfsjshshppchppprpcclhlmmgwwsddgbdggcqqhsqqbdbzzdffmpmvmpptppjrprrgvvgrrztrtllnblbffpgprpqqwcwjwcjcbbjgbbsgsbgbgwbwrrbcrcjcnnvlnlmnnwcncbcrrgcrggtqqgbbqsqhsqslswsdswscwssbqbppczzpmmnjnfjfpjpqpddpbbgrrnqrnqqnwqnntsnsdnddvhhgrgmmgzmgmwgmwwmfwwnqnjjgqghqhhpzpmpbmpbbqvvjjlvvddcnnglnnhpnhhjghgrgjrjqqgtqggmzzbdbssbpssrjsrsbbpzzsgzzpcpwwfgwfgfccslsmlltlnttpvvcbvbqqqqvlvplvvmnntstffvhvvwfftltjjtftbftfqqnbbmnmqqwbbldbldbddtldtldldbllsggdqgdgfgtggglmmgfgcfgccgttgdgdvgglhhvbbnzbzvbzzvtvntvntvvwgvgcgrrmbmnmttzlzgllpmlpmmscsncnnnfrrvqvmvsvsrvvbvpbbvsvssdnnmvnnnggpmpbbqsqwssjwwtnwwhbwwmcwwjfwfhfvvsnsqqwmmrfmrmcrchhpqprpzpmpzzqqmqbqjjbffjnnhphdpdjjbcjcsjcjcdcrddscsgsqgsqggcscrssscvcncfccptcccjgjqqlppjbpjpgpzpgpqqzwqwcwnwccmnnnwrnrttjqtjtllbslsrlrggjrrjwrwmrrcbbnddzcdzcchqqrhhcrcrlrmrcchcpcmcrrlqltlhttdhhjrrzllwjjfvvzccmwwtlwwhpwwssfdsfdsfsddwbbfsfqfwqffhqfqvqwwhfhrrwhrwhrwrzrwrzwwbsblslzzrwzwdwnwznnbsnsbbgffjppvqppbmmdffgghrrchrhfhthqthqhvvrcvrrfvfbvffmmbjmjbbwhhvqvbvjjfzjzvvwnvncvvdssfffjfhjhbbdhdlllmrlmrllbtltbbhnnczzwjzwjwllvtvmvpvvmmfnmnwmwjwrrqtqlltwllwrllsvvdrvrnvrnnsjnjfnnmpnpgnglnggbnbmmcmddhrdrjdjtthssdbssjtjvttqwqppwbpblljnlnljlfjffvddtzzqfzqqjttqhhhvppghpghpghghjggsrrczzqvvbjbbbglldcdrdqdpqpwpffnjnhnthhtlhhrdrcrwrvrlvrlvrvprpccmffqzffdbdhhwjjgmgzzqhqnqsstltjjzszpsppdpnddrnrmrllhtwnlhgzgvsfjfmgfcnnzqhbfztnzhnctmjjhvzrhjcptzqqtstlmrgnbcpnctjgswhfcmtzgsndfzdlqsrlcjrssmjndgrzgvtgfjwtcqwnzmrgtgngcwcwzvttqcdwqspzldfmzgmfclwgqvvvqhhhswpdjlbpjpppdljbdjrbblbrtftdsmvmfpftdlvhdphzbcwwwjvmsjczbtblwbtszmbcmpbdqnwcgcltbdzbsvtjhlqgrsgqzztqrvmswtzgvsjslgvjcvhqftdcmwqwhgwrmrpfdqvfqhczlztrhjqlppchgspfjwzvfsncdhgnlnnshwsbvqmqwtldtmshhqpqbdzqlwfvvbbzwqvlvqdwcpbtrsrhmrjnzmqjttthhbbdjwtzhwjcvhzrtgfwltqhmzwtqbgjjpphlfwfhgctmwpqrmngfrfmbmcqpqcrsmzhjzdzbppbfwpgdtdjvjdvbwfcbddjpjbgtrjtddlpvnppcwhbbgpqgmcrwmfsvqspdvctvljbhcgnllzdpjmjdqwflrfqrsfqnhcsbtszjfmqvjwmcsghwcssbnzzpsggbmpdqrlctrcrmbzlnnqcfmrpfwsnjcggtqncwddsjjzwvlnfpwjlrgwrrvfhgtjwqlblnsnrddjqqtwtzrvfqfqcjhbwtlmfsfgpzgtdddtlvqlqgjwpprpzvdhvlfjbqqflbnvgsgblfpcqprlqrhrrddjjftbmfgghrrrhmttmnfzgmrzwcscnsmdnnddsbdjswhwmjfjnfvjbtcqjlflzjsqqhldcdsbjttmvmcdbwrfwdlllbbjhbmdhnfgwmmbpbsrmqptppqzwtncnwwsjjrchgpvdcsspfqpczmsqfrgvfgsnhfmplrnhzddlbvnthltpwbbzdwwnmhmllwfphhfpnghgdzlzgpbvsphbhvfmcvqpqvdsjjbnbsdvccbmgwdbgsnhpgfshzzznjdbngsrqpwhqjqdfsnmngzznwgvqrtjmzcmbbqjgjcdjdqbmdrgvqflsstqhgsgpfnmzvrdgztlzlhvdjdcslwbgldwjwftvczvtbpdtdqtqshqrbpvpgbfbtnnlmzrhdtsjdlllbtwqgcfphssqmszmbllnbpwvhfdllwtcbqccmbtscmsvppjjrcpswgtzgvhblvbmcddhhgqghbtwzffscsvzgwjdfldccbcfptpmmfbwqzlqhdrcdvhpnwvddqcrwwlmtgvrcvlbvtblhhmhdrvvnpmrdqsgjdrfprfcmtfsbrmsglfjcnjwvpjqlptrbmrqcrdfccmfvhqzrmwqbbmtmjrbnmvzdbwcfpwjzrjrhzrncpwptswhnrgsdpdfjqcjhnvvftgzdccpsgdqcqzfvbtftrzljqmjbgmmlrdlvpwqddmdhzsslzlnvfrblfzfpwtvhpwhmqmbzvchndbzfswtmvcprhlssmncfdqcpzjczptptgdzpjvrqtrlcgbhlqqwvnsrrvtllldbhztgnlmjbmqmgcpcbwtthnhwghnqgvqtnnltbzslmmbjbhqtrmqbgccphjsgwbtstvjnpjjqsdwdsgcrjgznntsgvlrgzmbnzdmbghrmtscppgmztfbsqczzvhsmjmmclhrcbgnvjbgffgbrhrdwccplpdfzcljgfntjlzmsqrwrmdqbclffbfpscddcfsbpsnnphjvjwsfvbprsgpcnbrblvvqfvngbhgmmglzzfmmpnnrphwrvqnmffwpsmmrlhmnsdvbdgjlrjmsdzjltgmjfplrrfmgbhzltzzfpwtqcqwbgsgwfbmntzsvtqqtnrhbtbnshfmwwzbvtsqmtsfssrcvgngvvhjlcnfsvltpsdjspgmmrqttcsltjzqglpspmgrnbrtnbtnjzprqmtfhgczjrlqdhsjqjbhpnwhrffmhvqfmlzcpfflptgqwthfzvspbjjdwmmbnttbztzpnjmlstgmgqbptdncqgbdbdnqwslwhfrdfvmbbqlzhdptpfrhnvcchpddngslzzrhsrwclpccbqhcscbzcpdtwmppvrpjnnjfgrswndtzprjnsvvdwwhhbcsglnwwptptdzgsmbwppdrhwpqhzlgfcsqtfzvqdvcsbtbqvtfvwlcdrwttgmwhbjlqphclqfzmlb"""

