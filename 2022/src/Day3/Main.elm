module Day3.Main exposing (..)

import Browser
import Html exposing (Html, button, div, text, p)
import Html.Attributes exposing (style, class)
import Html.Events exposing (onClick)
import Dict
import List exposing (..)
import Set


main =
  Browser.sandbox { init = 0, update = update, view = view }

type Msg = Increment | Decrement

update msg model =
  case msg of
    Increment ->
      model + 1

    Decrement ->
      model - 1

view model =
  div [][
    div [class "sol", style "display" "flex"]
      <| part2 data
    , div [class "sol", style "display" "flex"]
      <| part2 dataR
    , div [class "sol", style "display" "flex"]
      <| part1 data
    , div [class "sol", style "display" "flex"]
      <| part1 dataR
  ]



part1 input =
    let
        lines = String.lines input
        totals = List.map total lines
        sum = List.foldl (+) 0 totals
    in
    (div [style "margin" "1rem"]
    [
      p [] [text "Solution" ]
      , p [] [text <| String.fromInt sum ]
    ]
    ::
    List.map
      (\t-> p [style "margin" "1rem"] [text (String.fromInt t)])
      totals
    )


total str =
  let
      lengthF = toFloat (String.length str)
      compartmentsL = Set.fromList <| String.toList <| String.left (round (lengthF / 2)) str
      compartmentsR = Set.fromList <| String.toList <| String.right (round (lengthF / 2)) str
      common = Set.toList <| Set.intersect  compartmentsL compartmentsR
  in
    case List.head common of
      Just val ->
        charToNumber val
      Nothing ->
        0


charToNumber char =
  let
      alph = String.toList "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
      r = List.range 1 52
      dict = Dict.fromList <| List.map2 Tuple.pair alph r
  in
    case Dict.get char dict of
        Just val ->
          val
        Nothing ->
          0

-- part 2 stuff


part2  input =
    let
        lines = String.lines input
        groups = split 3 lines
        totals = List.map totalG groups
        sum = List.foldl (+) 0 totals
    in
    (div [style "margin" "1rem"]
    [
      p [] [text "Solution" ]
      , p [] [text <| String.fromInt sum ]
    ]
    ::
    List.map
      (\t-> p [style "margin" "1rem"] [text (String.fromInt t)])
      totals
    )

split : Int -> List a -> List (List a)
split i list =
  case take i list of
    [] -> []
    listHead -> listHead :: split i (drop i list)


intersectM list =
  case list of
    [] -> Set.empty
    (a::[]) ->
      a
    (a::b) ->
      Set.intersect a (intersectM b)


totalG group =
  let
      charList =  Debug.log "group " <| List.map String.toList group
      sets = Debug.log "sets " (List.map Set.fromList charList)
      u = Set.toList <| intersectM sets
  in

    case List.head u of
      Just val ->
        charToNumber val
      Nothing ->
        0


data = """vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw"""

dataR = """DMwrszrfMzSSCpLpfCCn
RMvhZhQqlvhMvRtbvbcPclPlncddppLTdppd
tVMQhFtjjWmsFJsmsW
trRtvNhfJhSzzSTFVhQQZQhHGphP
CnLMBWLwDMgMcwwdngdHGPVTQGpTHZdGPGpd
LLDqcDgwqCMnLWqtvzrzbbtJqPjJ
wQQwHNQLmbWQbQRHwHNFBbwqPfjqlzRMGRqzpSfvPlzplM
nCtGCZZtsGsrtDMZpfMpSlMlvlZq
cJctJCgVJsCJnDTsCthGhGLwBWBbbQmLbgQLQQdWbbbQ
ZWnNlTNTnhhQQlDNdmmpwrrrqBwjwjZd
GzvlVRSfvMVMvGlSpdCCdjmfpmBCdsqB
bzlFlLzJWLHHttLL
SmzFhVDzrmSrszVDVhSVbhZcCZdfZNcnMfMbZnNN
PTTRGqwqTqWRwLgTLTZGnCbZbNddZCCtMcNs
sgPqPqgJgWWqjjwgwLLVFBFSVmvmBBrmJJDSvp
CBccfSBhcSBddfgtlJmmmwmPRwmh
FpTzzGWHWprgDtJlDZDPFR
HbbTzWnTrnWtCbQBbQqQbSjf
fPHspCjgwZggSvZQ
RrNhzFZFcNzFLNLNwQlSlLnv
TRFrcDVrrRmrhFRZzVrczqhRpjqjsssCpfHjsCdpsPfpjCMC
ZBnBTMVTSbGbTVTGbCPgqsgPtHtgCcPtBB
ldDrmnnNrzhdhfgcsHqcsfCcsCHg
zFdrzNdzQNDDhFdWldDrJTbZTbLZJVVVpMVWVnLS
pLnpQNhBbnWvbsWm
FrFwjlTPTdTqqrDZWbvmZbpSgmJWvbgS
RqDqRrdGFpGRFrFFdTNzCcHcHLHBzQCcNNGN
bvRCtbtCPfSGtCcvCbPNlglqgqlGZMzTgnlZnq
hrmWWFspsHWrzNwTnFlTMTwFFn
HpjJDWBQmmQdRffbPtSzdJ
GpHjFsjMFpCpMWjMGCqWmmqrWQtmthdbDbbD
fzgLTJwJPSgJgzSzPfhmqmQhQHzQtbQDrrmq
RNlBRwHfRJHLLfHTwLSBppNGvjNMFFCVpVFcvcFC
SfQnfSFHfnvMtQQSnHJtMffsdTlZtdZmtllmTlmlRRbBRLDb
hrwhWWwqzPrcCzwwzmPlRbdmlQDTPPBLDl
CpwCzrwGzNCWrJnMvpMvfJVFvQ
rCRPpgSggcpqrhPrCDDTLsMLDSDGLTMGVs
HdvzmRWmlHzwvWHvRHRvHJbDFsdMssQQVGMDTMDLVLLFLT
JBlBnnWBJlCqZRRqRBpr
GtZllZDlfDpGHZtZBGBZpDmzQzzCSVVFHmmmsPCQWWSS
JvFJJrwvNNcJTnbrTRNRSCzqwSsVPCPQqmCQszVm
JLMnTbLnMgBhDFDf
lffDhtgDJzCJNGGTzWTRRnRvBv
qpbpdwqZwqZSwMPSqdQcQmTRnWvnBnRBQBtVnTvWmB
SccbbwSbZbFPswpSZtgFlClLhgChhhNfJlFj
ClmCjCJBjwBVwJGjlGNFJlVMHSrfpDpTfrHMcHTppQVrHp
dRLZWLvWSHmTccWW
ggtqzmRZnmhghZhZghntdqsvBbjlbNFFBPwNNJNNCBlwPGBz
HZmsFZQpvZsWCZQvWrghGrhtgNzdHddHGh
fWSbqWDJVwcSccNzrNhcBtGcgG
VqVfTJnbWjqTSqbwDRfRvQvFpFLRpZsssQsCQZ
FpFZNfplSTJmbZzddGzhDrWh
LqLPPQjLLRMPqvjLLHQrLqrRWdzHnGhdthttGGbbDWhDDdWz
sMLMgvRLgscrLrRQvwmTNNfpNplglplfmp
MPVBmCmWGWRPPqRhLcnjcvjjcpjMvp
tzwrwsJlrldJsrsrTtrzrTtSNnLJSShnjcncvnvqnSFnqN
rswrzsbdwDHHbWZqVfWV
dVmmMTmBPTrCBRMCqFHSWFFHWzCvCz
jNqfGsDqtsjGjQjDlcJFFFznFtzznvtFpFFp
fNNhgsDcfNflqchVRdgVrRPRdVTRRb
HJPLwgLvjttmgHJFjwHgtlsBbNbbNsPpblspTllThT
MzmcRRrdDMVTzbhNNSszhl
mCDDVqdVcdDrqfcCnrFwtGwvngwvtWJtFjWW
dFDpmttBlvNNgWlglNDBFttmTGHTcSSJJHHnMsJsGGSdqcJj
zLbwMLVQbQRwVsJsSHSsHcJqbj
wfVZLPzfZZmpMZZMBl
PZHZMJSTBWHNWSHzVnhhfnhThhpnpC
jFdBBtrFjpfnjfnf
ccGrbblbGRDQMlvmQBvmBl
PCCTsnbPbHDnlDfDNB
rMjQltgSqtvMjScQggjfVVzBzFHzGfVGDLGBqB
vdtrMSttcdwcpSQSdglMrtWRRPJZCpsRZJmWRRWChWPh
pWzbsPCCPPpbptSMCJJwBQQGQt
cDDmcTTRRqzFRddVTSJwMShMtBwhwQMDwv
HldqVmVlZdLTcmRFdrngNNzrffjWpPLggP
JPqvjJmmqvSLmPtpZdcftdmtfdCC
swwhDRwBBHjFFBtBfZ
RRzNQDwznDsDwWJjLNlrSPLSTr
VQmdRLvDlmqZdFrBBJdW
CMstGsnHnHGGMrMZwMpwBSbW
GnsshssNfjtsnggnHCGhjtmfLQQczllvDRVTTQllQWlQ
dhbNbswbwVdNtVdstBtgbNQTBCCSFTmfmMFmfRqQmmQM
HFljLrvZfMQQQPvm
WrpznLZZrnplpWbgdFcFsNzszgst
LjddfTlMccnBfDQBtBQb
ZRSNchHwhNNCHNSWPQFFFHDBBtnQDH
CNpZshSZgpwJmpdLMlplMc
bTmTFmqzgbBntRVsFvVwcv
CZfMrlZLLLMlfPZRLRHGstnjGwtvGcsSVwtcSGvn
ClZpMLCRMZMrHMLmWpqQBpzpgQzmbg
jDmSSGWDDdWdSqqDDqCqpJzqRRqpJnRsMRcMzM
lPgNPvPrrgNrPhNszFggnRzccbMJgz
ZQTHQvQTZPrrQlBBrNvQZZGtTtGdsVCGsCTLLGDmLsjt
rbCfBrbsvQqRFZRNZC
HLSTcwqwZSQFFgRZ
wdDwjpMHqJDTMTdqjlfBvGBhsbfhbsnb
ZhZcHHHlhgchHhlCZZhLCCbGdrsBBGPNBjGbsjNNjnJnPn
wtJqqwDqQQMQFqSqTzwzVTBnGdsjBdnMdPGBBsBdnrjr
RVzJzmSVZmLLWpCc
gdqjQQrlhhQlQrhsnjjhLgmmvmHBBmTmZRsHJzTBHRJv
NwNnGNbGPbmTGpJzppBG
nDnVDfMDrQqQStgM
MLbgbppMMgLmHgQttGQJgJrBShwNShWBBSNNrwNqNN
GnTFlzCVVwPRrVWhSw
GnDDdvdnZDTdnGMsHbbttZgttLbc
mdmPmjClrTzqttfm
cpFnSbcQQsqNNtqWJzHS
QFpcMMBcZtLpQBjVjZhlPjjVlwvw
spVsPjTZZMpZMVLDjmdSQJCLJSmLzdJQdl
HhRnNrqwMhNhnqnHwGNRFBNBrzSCSdQmQCdddbrQSSclSSbQ
nFNqGRvqBfjMvTssfZ
FjjzjnpFqqzFFqgFSZjBhHfHhnHRDDwfdTdLfD
MmCMGCsMWbtJrtCWCbmsmWWhdLGGwRBwdfdLhdTLhHHTBd
bJmtrRvRjgzFFvVq
RWwWmVQGMFGmMGVCVWRRZSBgDdSdJGlSJpcBGGSlpJ
HHhQThnjBDHfSBlS
bPhNjbsssQzFNQqWmz
FTDtrjqwwqGtDbGnfBlcnLcWBZwlWn
mMhRMvJsJvJnMHCvHmhLZLQlhWQBBfPfLPBZ
HRCCsdNdvNmCvggFStbzjbGSSjjn
sLGddsvvcZmLvrLMGcMsVnTTlqlHCsTHHVVgVt
wRbfJPbpNRffRJMBhpDntTCHFNVgqllFlqggHC
DpfbPhRDJPMJppJwzfpbbDGSjrGZvdccQdjGvQZdvrLz
wTwLNLVTqnLMsBwfMsJmCj
JhlGvcdJhSFvFvvvMfgBpCzjzdCfsMMs
DSlPPJSGWrDcFPhtFhWJZHQZLTQVnRWRbHbZHQQT
TmTgTrPDNLNVlDrmlbgNmrSSGbzjZGMvjpZjvvphWMzW
QtBfdfQcdfHtZcnZnGZzchnp
HQHwRBGfBCGBtsrCNPDTmTlNLr
bfNhjhNJDWhlWhlRRR
SsscnHgnSnZnltqqfWRWrzZv
cnfTMfmMnTnmFGsnTVLLLpQJbpbbjpdTdN
BqwZzqRQQRRPSlFRQDDwdfWwhJphnfgfnpMdJfdM
rcTLrcrvDDChWJhfpTgTJh
DHGbGNVCZStGqSqS
dlfdRNNfVdLwrRnwdwRmhLFsbsJJgLqbgCBWBCsW
PHDppMPMHHDPzmBBCmWJqCmbJD
HzzZHmZzQcNdRRdZwddr
wrlshslPsSRfvrQvwbrslCDghtDgCVhDhBVJCFHddt
mZnGpWpWzGTMqnFqDqJNDNFJVJqH
LjpzGcjMGcTzcmmznWSRsfRPfcrbFQcfrwcv
rWBmmmtNmmtBbtlwnhJJVZbw
FsRcjGdLdvFslZbQJZwQps
GHFGvMccFPjgDNbmWWBTTHNz
GhHzmhmwlpbltGBmBmsZsBZsfCWC
rgrcCCPdsWBgNVBD
RnRMdQPMCqndSdQdcQhblpLLwhJbbpzGzwpS
NNQtStFPpJwhRbRzRbqpZZ
jLnmdJrrdDTdbgWbTbRW
JHvnMCmDnMnMljLCDmMLjHNFGGNBPVtQQFtSNFQtPQBv
BFbBRllFZJnPVJbV
GpGHwgzcLhDcwttwthzzhHcPTjZjMgMVZjgZTMmTnMZWJVJm
GccwhqcDtlrPqQrRNQ
gWHWLgHBHQdFhjGGThTQhR
pZsSMpZMJJSzMszzzqclpfjvrvvcRGGTcTVhbVvRbTGTRG
lnMwsqZqsslpjlSMSsffZqqJBgHNNPNDWdLLgdDgdLHPWwCw
qfNvBCBfBqfNMBqCZZfcnmnvtwScpwFSpSsSwt
HzdVzLWPPGGDdnsswnztsRsnmn
QddWVQgJPPHJTJbjBtNTTq
DdRDDPRGGPGccfcbJwsbJWzsnznlLLWzWTLWhVVVVS
CvCrNCqgFqvmqNZFZqqZvpWlVrlVhlhnTLShlDWnzVBD
jvqpvpvpQNCQQCZZmmNgZfdGddRjJDPRMHcHJDHPJf
ttdtBtPPMqWMdgPPBbVGWfTGTTzSVLfVrzCS
ZpDpvRpZDDcmmjmZfLSrwzRnSVSnwTTR
ZvQmjFVHJFDcQjDlZcDVHdqMNtqNBPtPJtbhhbdbts
dGdwwLLpgwgssJpgssNhpJlnbfjnzFfcbfttGjzjlntf
VQvDvHVVQHrQHDCZVBChrHFtzffnfltFFtncnvFtllMl
VBShSqDVRVSTmppPwwsP
fTFDTLNNzlcNrmDcrMDTFPwCSsbCbPPsnCPwLSPvbs
ttQqhJtBRRGnvgHGnlSnbl
hZBJlQBRjVRBRjhtRRMNFVmFmfDNrfWcFVmD
mcTZFBFmqBjmBgPtCtPprmssStCP
LWDQNqDJfQNJddnWfzhfsPRVppVVsSptftpVMS
NDGnJDDDbzddWdNbGNQQLQbqqFBBFcjlZBlHjlZHGBTvZB
PwDzvphPwVwWBqLLwnJWTq
jdCGCgjmllCrmmlmjrbgmRdgJSSJJFLSSqJfLnqLLLbWffLB
mRdjcMHgDpZhDqMZ
cqLjhhrwZwJbBqZhMwbZZdGWdGSllWFvLFGQdnGFQG
gHHVzzppRVggcgpcGWRQRSvdSvvGWvll
HmNNHtVggHsHPtrhJsbjbwCrCqJc
zqPvzLVvzFFQZzWpRLlmHRDHmRCHDH
dNjnJGGrGdqqMprRlpqB
GsgtjhSsSvvSFqvP
pVrfzzjrjWVWTWjrNZvnJSJZqnnqnpSZZS
bdQVQPRPDdcbRGPFddRVMVlZlMlBqSBBZSvSZwnwvJBS
bFbcFbCPPCbbVHCCdVgWfrzjmWfrWrNWgHfT
JgJqLjjjVGgdqGDZGzlGRStStT
PHrHccmrMrTSMVStRtRR
HWPWffNsrppfPWNsVFsmPNCJwwjdJdvdvnJwghBLJLpdLJ
HtHvcnDSDgDcDHtpLrvwjwjfZMjffw
CPWzdJdqVdWZpnLdwnrfdn
GNCNmTQnPVRRglSlHsSG
FJdhjTPbdPJjTPjTjPtSLsSBWWRcCvCvsBWztc
MfGgrHMDDpMnZGDLCRLScCsBlgWvzB
HnmpmNNHGZZpZZrnMPFFbNCNbFdTPVFFFN
TJrrrJQTqJqmTltfRrgfgtgFFg
jLRzBvBjjcnFBNwWlgBZFt
RMjMCGpGzGznzhRmmPPDPsmMmPQmJs
BZqwQCQZGZcVBczqBHtfbbbWfTqNWfMfPNqW
LLpmFjpvpHrvRFSRDRMWbdbtfPWPbjtMgMtW
SDnrpDprDFnQhZCVnhcH
WTsBBQTfQQTTbJBbZbnfTsRFwFrjwjFlrRqvrrlqvWRV
pGcShcGSLNJNHCLttlpllRFgpRFlRpgRrg
GzcMLScSGJGtCbsbQfbZbMBnBn
NGCLGjVjZjQwTGJRQdWM
cFTcvSrFmnnpSmndMswsRMJWRwMHps
rrrhhcTznqvzmcccvvmhgzqDgbgttlDtjjjlfVCfZCjZZV
ccDMHddWNDnnNWMMzdHJJmSQhfQZfvQZflrZQfdVfLLZ
bgBFRTwFtgqCgpRGFpvpVllZlhjrrlVlvj
wtbBGPTPtRTgbCTBqFgGRwFnsWJnmDMsWMJJMzHPhDmJzP
zsbsMtMMdnffBbzNsBtCCWLpLrCrcNLVDWVVcD
TmPhJRvwmjmhFJwjjRPFPTvJGVCcCGBrDpccpDrCrWCVDVFZ
QvSTvBhqwjPmwddHgtqMnllzMl
gftDtqnpqzGZsFcthbtZ
VlNPrBrRNrLBmdRVFCcGCZTFCsTCsbLL
VdldlljlSNHBsSlqfgqMDDvzpHJHWg
tQDLvFLcDrWrcnsHffCGgGHG
ZRPTPJqhMZJZVllRZJPVZPRHnhCnfdssnCznzGhdgfwCHn
qPqlPVlTlSqbZZVJplqlPDmrjWFtmLtFWgQvtmtFvp
zlZzdNRPgGGzsLGCDBBtCDCtSncScP
vWvHWbqjrFMbvrTWcVnQBBBSjLDcQJcL
wfLHwfFqLFbhHvWhMWqwbwwRspssmzgpzGgmmNfmzmRGRz
rPvLrQBvBLsLLdtrgssgZjwFwlnCFMtMFnlllnnb
mNmmzpWHlzjlJMJb
TVSVTWpqRWpSTqNbTVRBPDfLLPrSLrsfQrrvsf
nRjpQWnQnRQzMjRdrtvvPCfmvGtPfMcCtG
TDbrbhNZVbbbbwhDZDhbTTGfcftqcGVvmmcqcJCcCPmJ
NLhrSwgwgnsLsQWljW
JWqVSpGNPdNNzdZJJpMzHzwLgsMwzwQwMBgL
clrlcvrRfccCtFbHrBWLgwLHmMHsHg
DbfDFjcvRcvchWZVWdNpGZNqdh
sdfvFLfmtszQwLfddRpmtDDBjVNWGMNQVQNMJGWJMj
lccrncTZhqqcqhWggvrjMNMGrJMG
SblShnZCqSbPhhbcbTTSZFdFsFpmdRwPwzvmswLtmm
PGwwHpfnFSvVpWqWCQNNjCbbnW
lmddlhcDRBlLRchdmzbNjqqWTcbNPNWTzz
RBMrRdRhlDtPrJtfwFHpsvrHpFSrFw
hhwlglFFSQndLRFbmCbTTz
NczHMMqzpzPcpfBffcmTrdfGbbRbGrdGrLCL
qNzNPqMjcqNBWWccBHsZPDhJnllwnwvJvQnJhQsgvD
mbmvmvbbprZmlFmZbFgLffgQtFNHNhfqQtNQ
SJcdzjSJBzdBdJDzQhhLQfqzNQQHggRL
jwDwcTTDThvTZPPW
FSVBBBvHvCpVVDDGcGwNNhhctwMvMc
fLLZsZVQmjfTfqQRmQhhtgbbJbGJRghtcGct
qTsTQdqjVfqdVdZZqVLpCpzSpdppBlSpCFdHSC
sQQhWsMmQshlhmMQZFDHDJFjgjzHZgcHdH
LnwnpNRrnrbCqqLpwnqfnLcvFHJFzNcHzJcgJJHgdDgN
wCbnpCfPCVqwwnrrbbPRGMMlSllmlTTmsThVMlsd
pzrprfwgbwtwqzrCWbqCwqSMvddHdDSvtHRlDnRRDddD
zQLzQQjPBPFcLcQFTFsmNQzcMNdDdvnldHHvdvnDnRnlvRnJ
cTZGzzscLcPrqrfrZqqbVV
DcSdcTwDLmcwDwvWssGfJfcJQZPGnfcs
FlHFMgtgNggpsztMHMqpjgBBnCfPflfQnZCQBBCnRPZC
gpVjqNVrHFtjqqzSLDTSmTDwwrmhbs
MLMzJTsZzZMgMLgHMmVmdCVhCBlQwDwwhChD
vtPRQpbqCldwdtBC
bQqFbnQbcFfjPRFPQnTrMMgcJgJrssrzgrgS
mtdGJmQRFmdtQvdvtRtdHzHzqZqpHFzZnCzhZjjH
fPwVlllswMVNPfBDDlNVsMsfcBjchHncqzjZbpzjcqCnpHHn
rlsNPWNlhWTPMMNPfwNWTLQRvQLLmgvSJvRJgTRG
TwnqhqqgvQnGBGmBDp
SMjclJSjjVJgCzCzNgpmdBpmBGspRBmpDDVB
JjMCgMMHMMZNStllZSNHhPqFhFWfqPPqTqhLFqtL
lRQPtjPRlDdStDSlPmvllvLsCphFfCHLHggspgFmsFLH
qwpTNprcbNWVHLrfFssBgFCM
NTWTnzTTWGZZZVRSRRQGpdDtSQRp
gpwTPNPBPTdLLLLVGl
jSHdjzZHMcDVtDvFjtCF
HqfZMHzbcqRRRWgdqPmBBBNmwW
PvSBtdFgvSmBPngFBTBjbSjwwpGjsppMjNpMjj
VZLfVQLzQQQhllpcNcwbssvwwwZj
vHWLVVqWTmTgttgq
CNRmNRFNRCWbWNCrlmfGlWqFLsDZQZBZzgwQZsBsDZZCzczB
MSjdVHvHnDDhHvdwBwssZVzwcgLcQg
HnMMTdttHSHSpHvDddpSHTjWlNWFlmRtRmRbqGfqGGNNfR
fBLTDppznrfTndfnfTzTLPvZvvHVbRbggjvzVbzvbV
mwmDGGlqDhMqthGqhJMWmlNVRZPHjgwjjRZbbHRgRHvv
DmhsJsshWGhSGlmlmrdcLLsTBBfcfnBppc
mbCGFFmGmcdTrCTQdh
MJHfJNLllJffPLRTdBqTRQNcqQGB
fPJHfSSSWfSLDMLWGHDMLDFmznmsjmvZwzvjZjbvbZ
pPvpJSfZTTvCzNZczzQZchcj
svbHWsqsvbsMFtVHgVtcRQcDlQRRRQLjlqjczj
tBsgvHVMFggbgFrgWnwSndfBmmBJfPSfpn
jwbwfjSbwjVSjvZPzWSvhvhQlCsBFgLRLLgBLRClLLFQQw
GdNJHpmHTDnTNJqnFCgBLFLFzFtsQRCd
NpMJHpnMrDpJGTHqTTmJHTPjfcvbWfrffVzvZfVWSbjz
wFwpqWwwpqwtqqrbCFtptDmCcfNhNRzRBZRRJRChVNBZBJ
svlvjHsQlvdlvMLdlvPSSLtzzczcNhJthfNtRcNMJNMc
HvvPLSHjgltjsvqwbbnmWmDpgwTT
zhCmPVwwChdCBtsWnNWswBWr
GJJSfSgFpjJjGgpfpgrcNNstvnBHNnHLtFHr
jgDTfjpMgZMGMGJTMMJRhzZPCzbhVlPqdNCbhd
bDbQQmVDRpDNbRQlfQQZnfwTlllfsT
FChzzBWhVzrgMwffJwlnngnTlJ
MCvqvhFzcHCChjtpNNVLppGmbq
bZZzJnccqdzcLhrcQDLrDs
FfCfWVfjWTFClClfwjWCfGGwhZSDhSLsSSRpZprLph
mFmTMmFjMMWFfZtttflWjmWTngNHJHggJJHtzgnJvBtBgHdv"""

