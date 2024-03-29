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
    div [class "sol", style "display" "flex"]
    --   <| part2 data1
    -- , div [class "sol", style "display" "flex"]
    --   <| part2 dataR
    -- , div [class "sol", style "display" "flex"]
      <| part1 data1
    -- , div [class "sol", style "display" "flex"]
    --   <| part1 data2
    -- , div [class "sol", style "display" "flex"]
    --   <| part1 data3
    , div [class "sol", style "display" "flex"]
      <| part1 dataR
  ]

part1 input =
  [p [] [text "sol part 1: "]
  , p [] [text <| String.fromInt <| processInput input]]


processInput : String -> Int
processInput input =
  let
    lines = String.lines input
    fSizes = Result "" Dict.empty Dict.empty
    d = Debug.log "result: " <| processLines fSizes lines
    e = List.filter ((>) 100000) <| Dict.values <| .folderSizes <| folderSizeFromTree d    
  in
    List.sum e

folderSizeFromTree : Result -> Result
folderSizeFromTree result =
  let
      children = List.reverse <| Dict.keys result.treeChildParent
  in
    Debug.log "with sizes: " <| addSize result children
    

addSize : Result -> List String -> Result
addSize result children =
  case children of
      (f::r) ->
        let
            size = getSize f result.folderSizes
            parent =  Maybe.withDefault "" <| Dict.get f result.treeChildParent
            parentSize = Maybe.withDefault 0 <| Dict.get parent result.folderSizes
            sizes = Dict.insert parent (parentSize + size) result.folderSizes
            newResult = {result | folderSizes = sizes}
        in
          addSize newResult r
      [] -> result

type alias Result = {
  current : String
  , treeChildParent : Dict.Dict String String
  , folderSizes : Dict.Dict String Int
  }



isCommand s = String.startsWith "$" s

isCdDir : String -> (Bool, String)
isCdDir s = 
  let
    list = List.reverse <| String.words s
  in
    (List.length list == 3, Maybe.withDefault "" (List.head list) )

type Content =
  File Int
  | Dir String
  | None

isFile : String -> Maybe Int
isFile s = 
  let
      w = String.words s
  in
    String.toInt <| Maybe.withDefault "" <| List.head w
  
isContent s = 
  let
    isD = String.startsWith "dir" s 
    w = String.words s
  in
    if isD then
      let
          wr = List.reverse w
          maybeDir =  List.head wr
      in
        case maybeDir of
          Just i -> Dir i
          Nothing -> None      
    else
      let
          
        maybeInt = String.toInt <| Maybe.withDefault "" <| List.head w
      in
        case maybeInt of
          Just i -> File i
          Nothing -> None      


processLines result lines =
    case lines of
      (first:: rest) ->
        if isCommand first then
          case isCdDir first of   
            (_, "..") ->   -- cd .. don't need move on             
              processLines result rest
            (True, dir) -> -- cd xx store xx in current
              let 
                newResult = {result | current = dir}
              in 
              processLines newResult rest
            (False, _) -> -- ls     don't need move on
              processLines result rest
        else -- not a command so file or dir
          case isContent first of
              Dir dir -> 
                let
                  newResult = addDirToTree dir result
                in
                processLines newResult rest
              File size ->
                let
                  newResult = addSizeToTree size result
                in
                processLines newResult rest
              _ ->
                processLines result rest
      ([]) ->
        result

addDirToTree : String -> Result -> Result
addDirToTree dir result = 
  let
      newtree = Dict.insert dir result.current result.treeChildParent

  in
    {result | treeChildParent = newtree }

getSize dir sizesD =
  Maybe.withDefault 0 <| Dict.get dir sizesD

addSizeToTree : Int -> Result -> Result
addSizeToTree size result =
  let
    oldSize = getSize result.current result.folderSizes
    newFolder = Dict.insert result.current (oldSize + size) result.folderSizes
  in
    {result | folderSizes = newFolder }


sumOfSmall dict =
  0

-- when cd update current folder
-- when ls next until $ are content of current
-- when dir build tree of current
-- when number add size of current

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


dataR = """$ cd /
$ ls
dir fts
dir jnwr
dir lrvl
dir nzgprvw
dir snwqjgj
16394 tllvcdr.sjl
195891 zbdp.gqb
dir zddrb
$ cd fts
$ ls
dir dlqtffw
dir rbfmmjvd
254713 wvwhrb.dhh
$ cd dlqtffw
$ ls
73533 nqbvg.fgd
$ cd ..
$ cd rbfmmjvd
$ ls
290697 zcgrgff.fnf
$ cd ..
$ cd ..
$ cd jnwr
$ ls
323577 ghmtnzr
57588 tdcbdpnr
dir wbv
dir zzbvdcf
$ cd wbv
$ ls
dir nmdwbnnr
253584 slzdbm.ncn
$ cd nmdwbnnr
$ ls
208370 scbcsb.pjg
$ cd ..
$ cd ..
$ cd zzbvdcf
$ ls
8052 ssssrhwz
$ cd ..
$ cd ..
$ cd lrvl
$ ls
dir bqqltcg
189288 cwpwh
90813 jhnddzml.lww
dir pwc
dir rjl
dir rzvqvv
dir slzdbm
dir tcbjmq
140665 zbdp.gqb
dir zhbpqlnd
$ cd bqqltcg
$ ls
dir dlbjblbf
dir fdlw
dir jnwr
dir slzdbm
dir zcgrgff
$ cd dlbjblbf
$ ls
11732 rnsrrf.rtv
$ cd ..
$ cd fdlw
$ ls
dir hlvpfw
228436 mzsfcvgv.pqw
$ cd hlvpfw
$ ls
dir dhwq
$ cd dhwq
$ ls
127459 cgdpqh.tct
58310 jnwr.nqt
$ cd ..
$ cd ..
$ cd ..
$ cd jnwr
$ ls
305998 ssssrhwz
129135 vrt.qdt
86204 wnvm.bld
$ cd ..
$ cd slzdbm
$ ls
40915 zbdp.gqb
120991 zsvffwlt.rbp
$ cd ..
$ cd zcgrgff
$ ls
94614 jnwr.mqm
191626 zbdp.gqb
dir ztrslh
$ cd ztrslh
$ ls
dir bhzn
201167 dvcjtzvs.rvd
dir slzdbm
dir szrth
dir zcp
$ cd bhzn
$ ls
119164 qbgmrqw
102090 zbdp.gqb
$ cd ..
$ cd slzdbm
$ ls
124928 gtdq
$ cd ..
$ cd szrth
$ ls
dir hpbbsq
dir vlwlsdjp
dir zcgrgff
$ cd hpbbsq
$ ls
151717 qsdhff.jsv
$ cd ..
$ cd vlwlsdjp
$ ls
15049 glpdjtq.mwm
302526 jnwr.tds
9550 lhwtbh
22857 ssssrhwz
$ cd ..
$ cd zcgrgff
$ ls
73640 mpq.cdn
dir zcgrgff
$ cd zcgrgff
$ ls
115955 rssmrfbq.trs
$ cd ..
$ cd ..
$ cd ..
$ cd zcp
$ ls
dir qdjtmwrr
dir wpdjttm
$ cd qdjtmwrr
$ ls
138185 jnwr
$ cd ..
$ cd wpdjttm
$ ls
207755 vvwtghjb
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd pwc
$ ls
15911 fqw
$ cd ..
$ cd rjl
$ ls
119274 ssssrhwz
$ cd ..
$ cd rzvqvv
$ ls
273935 ssssrhwz
$ cd ..
$ cd slzdbm
$ ls
dir hlvpfw
290414 lgzbzjvd.glj
dir ljpmphn
316440 slzdbm.tsj
$ cd hlvpfw
$ ls
dir mhvmszgh
107004 slzdbm
$ cd mhvmszgh
$ ls
dir tftstdp
$ cd tftstdp
$ ls
176794 mpq.cdn
$ cd ..
$ cd ..
$ cd ..
$ cd ljpmphn
$ ls
37528 slfdb.bqt
$ cd ..
$ cd ..
$ cd tcbjmq
$ ls
96506 lrz.nhb
$ cd ..
$ cd zhbpqlnd
$ ls
14027 ccswrthv.pfd
dir clnqtjz
dir fdsqmnn
dir fwdt
dir nljb
dir npmsdrgp
57812 slfdb.bqt
184299 wmlmg
241025 zcgrgff.wbh
dir zggqfj
$ cd clnqtjz
$ ls
62391 cbhw.wgr
309318 jlm.lfq
dir mzsfcvgv
dir rrn
307583 ssssrhwz
$ cd mzsfcvgv
$ ls
1356 hrbh.wpz
dir vnwqstw
$ cd vnwqstw
$ ls
184434 hnhzdshl.lrl
150624 wsrnprdb.pjf
$ cd ..
$ cd ..
$ cd rrn
$ ls
105792 dzprqqc.rbh
107482 ffdjdbc
dir hnr
dir rdmgtf
dir rgrwp
325147 shqr.pdj
43186 slfdb.bqt
236667 zcgrgff
$ cd hnr
$ ls
dir gljrlhjm
250526 mzsfcvgv.nsb
43164 ssssrhwz
$ cd gljrlhjm
$ ls
dir wjwqrnj
$ cd wjwqrnj
$ ls
142366 gshl.qfc
$ cd ..
$ cd ..
$ cd ..
$ cd rdmgtf
$ ls
193562 gdrdnc.vml
123723 hmqfdht
dir lzfntb
dir mjfwsmgd
208819 nqbgcn.qfq
dir tqh
$ cd lzfntb
$ ls
dir gwnpsvw
dir rsgwzhp
103487 scvllbjh.pnw
$ cd gwnpsvw
$ ls
dir lqnz
81937 zbdp.gqb
$ cd lqnz
$ ls
27250 zcgrgff
$ cd ..
$ cd ..
$ cd rsgwzhp
$ ls
dir hhz
$ cd hhz
$ ls
46435 djvfz
$ cd ..
$ cd ..
$ cd ..
$ cd mjfwsmgd
$ ls
25726 clcvm
170085 zbdp.gqb
$ cd ..
$ cd tqh
$ ls
111272 gdq.llg
3215 hdghs
dir lpdcdr
dir vhfr
$ cd lpdcdr
$ ls
203902 hmqfdht
$ cd ..
$ cd vhfr
$ ls
91584 hmqfdht
dir stmvj
$ cd stmvj
$ ls
315660 wwpq.ffq
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd rgrwp
$ ls
dir dhwtfld
187439 hmqfdht
dir jnwr
dir npr
dir sdsppq
228581 twd.nnc
42349 wbf.cwb
194162 zbdp.gqb
224441 zcgrgff.qpg
$ cd dhwtfld
$ ls
152381 dqdrd
$ cd ..
$ cd jnwr
$ ls
229758 hmvw.gdz
92710 mzsfcvgv.cdd
79954 stjtn.gft
89831 zbdp.gqb
$ cd ..
$ cd npr
$ ls
52767 hlvpfw.rqb
dir nrbjzvgq
270579 pbsq.msg
240181 slfdb.bqt
dir znbwnh
$ cd nrbjzvgq
$ ls
295237 bwqqvn
229608 dhrjnvtt.lvm
55391 nlvn
dir zcgrgff
$ cd zcgrgff
$ ls
124604 hlvpfw.mlw
$ cd ..
$ cd ..
$ cd znbwnh
$ ls
228293 zcgrgff
$ cd ..
$ cd ..
$ cd sdsppq
$ ls
235741 bcsdzpfj.lvd
dir jnwr
dir njtrhrfm
dir tvq
dir wshgn
$ cd jnwr
$ ls
273541 brcps
dir gjt
dir hlvpfw
dir nbsdvpnj
dir zcgrgff
$ cd gjt
$ ls
184707 zbdp.gqb
$ cd ..
$ cd hlvpfw
$ ls
242810 zcgrgff
$ cd ..
$ cd nbsdvpnj
$ ls
181797 hlvpfw.gsd
294284 hmqfdht
215098 mqbclwwq
$ cd ..
$ cd zcgrgff
$ ls
188814 hmqfdht
$ cd ..
$ cd ..
$ cd njtrhrfm
$ ls
dir dqpjdztt
$ cd dqpjdztt
$ ls
178036 vwg
$ cd ..
$ cd ..
$ cd tvq
$ ls
233085 rmq.zgq
215613 slzdbm.wvf
$ cd ..
$ cd wshgn
$ ls
209007 jnwr.hsr
dir ntrlll
$ cd ntrlll
$ ls
245558 jnwr.tbr
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd fdsqmnn
$ ls
dir dfl
dir hqffnmq
268076 zcgrgff.cfm
dir ztc
$ cd dfl
$ ls
5553 bbvqpgf
128390 lqtvvc.tdj
190189 vnmfcdws
$ cd ..
$ cd hqffnmq
$ ls
28438 ssssrhwz
$ cd ..
$ cd ztc
$ ls
247470 slfdb.bqt
$ cd ..
$ cd ..
$ cd fwdt
$ ls
60342 bdwbzpzz
dir cdcc
dir jldlzttm
261042 ssssrhwz
dir sws
293090 zbdp.gqb
93624 zcgrgff.vgh
$ cd cdcc
$ ls
dir cwpfczr
dir lsrgfml
$ cd cwpfczr
$ ls
303853 wzz.chs
$ cd ..
$ cd lsrgfml
$ ls
63205 hmqfdht
$ cd ..
$ cd ..
$ cd jldlzttm
$ ls
184715 brclwptp
309967 gdbfh.vhn
5460 hlvpfw.fwb
46159 jnwr.ggw
173896 zbdp.gqb
$ cd ..
$ cd sws
$ ls
dir hlvpfw
75133 nqrp.jdg
$ cd hlvpfw
$ ls
71403 cvdd.whl
$ cd ..
$ cd ..
$ cd ..
$ cd nljb
$ ls
105380 lrhmdl
$ cd ..
$ cd npmsdrgp
$ ls
176062 bhm.gfd
dir dgdl
24535 fdfmntpt.qvp
dir jnwr
279503 ssssrhwz
4671 zbdp.gqb
$ cd dgdl
$ ls
4624 qqcp.rcg
$ cd ..
$ cd jnwr
$ ls
dir cvw
122648 hmqfdht
26565 jnwr.fst
271544 jnwr.tqb
170968 mpq.cdn
dir ncqflwrb
103826 qgzlff.frl
dir wclwg
dir wtcfswt
$ cd cvw
$ ls
58583 pjthqdbr.zzh
398 ssssrhwz
$ cd ..
$ cd ncqflwrb
$ ls
105747 wwstfng.cdl
$ cd ..
$ cd wclwg
$ ls
75020 wbq
$ cd ..
$ cd wtcfswt
$ ls
182813 dnln
255923 hmqfdht
$ cd ..
$ cd ..
$ cd ..
$ cd zggqfj
$ ls
dir gpnfpv
dir gzfsnbv
186684 jnwr
dir msrz
$ cd gpnfpv
$ ls
dir hlvpfw
$ cd hlvpfw
$ ls
186814 bgqpn
41444 mzsfcvgv.qpp
292879 zcgrgff
$ cd ..
$ cd ..
$ cd gzfsnbv
$ ls
321796 jnwr.ghb
dir sdgwjtf
$ cd sdgwjtf
$ ls
127216 mzsfcvgv.qlv
$ cd ..
$ cd ..
$ cd msrz
$ ls
dir fzl
32371 jqtgpzw.thg
271389 slfdb.bqt
270845 tchtnp.rsw
146371 zmwbmj
$ cd fzl
$ ls
140610 ssssrhwz
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd nzgprvw
$ ls
dir hgfdbf
218138 mpq.cdn
73419 slfdb.bqt
dir vfztvnjm
269866 wldptfv.bbb
dir zgdsgh
$ cd hgfdbf
$ ls
163792 ssssrhwz
262305 zcgrgff
$ cd ..
$ cd vfztvnjm
$ ls
139863 hlvpfw
138410 hmqfdht
$ cd ..
$ cd zgdsgh
$ ls
194820 dhp.tgt
dir hlvpfw
dir sdllphqd
dir zcgrgff
$ cd hlvpfw
$ ls
44509 hmqfdht
$ cd ..
$ cd sdllphqd
$ ls
47133 mpq.cdn
$ cd ..
$ cd zcgrgff
$ ls
186613 jrrpljrc
$ cd ..
$ cd ..
$ cd ..
$ cd snwqjgj
$ ls
dir slzdbm
dir zcgrgff
dir zldscsc
$ cd slzdbm
$ ls
21422 dplgjzvs.nrn
175310 zcgrgff.dfq
$ cd ..
$ cd zcgrgff
$ ls
dir rrfhvjf
275455 vwnvj.nhb
dir zcgrgff
$ cd rrfhvjf
$ ls
69149 mqmwv.jmr
$ cd ..
$ cd zcgrgff
$ ls
dir zcgrgff
$ cd zcgrgff
$ ls
183398 hmqfdht
93968 sdhmwd
$ cd ..
$ cd ..
$ cd ..
$ cd zldscsc
$ ls
dir fmzvccvg
299187 mpq.cdn
324288 mzsfcvgv.srz
dir prdfldrf
dir sgpdqw
289150 ssssrhwz
3330 vwwj
133537 wbb.fdl
dir wjdpws
$ cd fmzvccvg
$ ls
dir brrlg
dir crm
dir mzsfcvgv
$ cd brrlg
$ ls
17776 jnwr.qqz
$ cd ..
$ cd crm
$ ls
220545 mltrj
177044 mzsfcvgv.thj
$ cd ..
$ cd mzsfcvgv
$ ls
51223 pdnbml.qpg
$ cd ..
$ cd ..
$ cd prdfldrf
$ ls
213872 ftw
$ cd ..
$ cd sgpdqw
$ ls
205684 whrfvw.dph
$ cd ..
$ cd wjdpws
$ ls
dir hlvpfw
$ cd hlvpfw
$ ls
324522 zbdp.gqb
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd zddrb
$ ls
238213 dhrdjfzl.npj
28938 hlvpfw
313900 hlvpfw.bwd
dir pgbgpn
132967 rrhsr
dir slzdbm
142210 wlvpwjjz
dir zdhcp
$ cd pgbgpn
$ ls
dir jpmvrvt
154268 mpq.cdn
dir mzsfcvgv
dir vvmwgngv
$ cd jpmvrvt
$ ls
150157 ghcvqm
$ cd ..
$ cd mzsfcvgv
$ ls
93810 mjglpwbc.tvt
dir slzdbm
$ cd slzdbm
$ ls
dir jnwr
174522 slzdbm.ncl
$ cd jnwr
$ ls
183686 rpwjvc
$ cd ..
$ cd ..
$ cd ..
$ cd vvmwgngv
$ ls
dir bqc
dir mzsfcvgv
dir qgj
dir wfjzls
$ cd bqc
$ ls
191442 bqsn
dir lmnvtg
278861 mpq.cdn
191779 mthgd.mjp
dir mzsfcvgv
dir twzq
$ cd lmnvtg
$ ls
dir dzdd
7920 hmqfdht
dir rnlctfm
dir slzdbm
213184 zcgrgff.vln
$ cd dzdd
$ ls
dir jnwr
$ cd jnwr
$ ls
71772 hlvpfw
$ cd ..
$ cd ..
$ cd rnlctfm
$ ls
282289 jvnmblr.mdc
24454 ssssrhwz
$ cd ..
$ cd slzdbm
$ ls
316131 mzsfcvgv.czz
$ cd ..
$ cd ..
$ cd mzsfcvgv
$ ls
dir nsvn
dir plj
dir tvmv
$ cd nsvn
$ ls
dir bcmd
dir jcjsp
dir jnwr
$ cd bcmd
$ ls
206778 cpljbvm.vws
86096 lfgjnzhw.rtm
$ cd ..
$ cd jcjsp
$ ls
188389 zbdp.gqb
$ cd ..
$ cd jnwr
$ ls
293633 jhh.wtj
$ cd ..
$ cd ..
$ cd plj
$ ls
323119 qgs
221100 ssssrhwz
264549 zbdp.gqb
$ cd ..
$ cd tvmv
$ ls
198197 mpq.cdn
$ cd ..
$ cd ..
$ cd twzq
$ ls
dir hwgndwpj
$ cd hwgndwpj
$ ls
dir srdzqqf
$ cd srdzqqf
$ ls
101738 slzdbm
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd mzsfcvgv
$ ls
dir hlvpfw
dir vbdw
$ cd hlvpfw
$ ls
270309 wncrt
$ cd ..
$ cd vbdw
$ ls
dir csw
$ cd csw
$ ls
231750 dcjr.jgb
297504 mpq.cdn
215564 plm.rtl
$ cd ..
$ cd ..
$ cd ..
$ cd qgj
$ ls
192181 hmqfdht
dir lmp
250495 mpq.cdn
dir slzdbm
253262 ssssrhwz
dir tzth
$ cd lmp
$ ls
dir jnwr
$ cd jnwr
$ ls
284951 zcgrgff.slq
$ cd ..
$ cd ..
$ cd slzdbm
$ ls
266715 hzbq.plc
dir jnwr
11379 mzsfcvgv.rbw
dir slzdbm
$ cd jnwr
$ ls
108636 mpq.cdn
324035 zzvmlrj.rrl
$ cd ..
$ cd slzdbm
$ ls
68393 zbdp.gqb
$ cd ..
$ cd ..
$ cd tzth
$ ls
12884 hlvpfw.gwl
146470 jnwr
$ cd ..
$ cd ..
$ cd wfjzls
$ ls
191088 hlvpfw.prw
64712 mpq.cdn
103828 pjhg.qmc
dir qvmdt
$ cd qvmdt
$ ls
244001 ssssrhwz
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd slzdbm
$ ls
dir hsvbbbp
dir jdtrdgqs
dir lrldsqcr
194137 mzsfcvgv.hnv
dir nzcbqzdz
$ cd hsvbbbp
$ ls
104752 rcfpnh.vlr
185473 slzdbm.dbr
dir vpwf
dir zhngz
$ cd vpwf
$ ls
266988 frhvp.ltf
82931 mpq.cdn
$ cd ..
$ cd zhngz
$ ls
181282 drtfhrhp
$ cd ..
$ cd ..
$ cd jdtrdgqs
$ ls
37106 hdjl.rrr
$ cd ..
$ cd lrldsqcr
$ ls
dir brwcprn
dir hlvpfw
dir jnwr
dir slzdbm
dir wrmr
dir zjrdgf
dir zrpqf
$ cd brwcprn
$ ls
dir bmjwts
$ cd bmjwts
$ ls
157148 jwbd
251396 qtgc
$ cd ..
$ cd ..
$ cd hlvpfw
$ ls
121723 zbdp.gqb
$ cd ..
$ cd jnwr
$ ls
240747 gfrrtdh
$ cd ..
$ cd slzdbm
$ ls
176832 jnwr
36125 ltwmzw.fdl
132059 plqmwhr.hhv
85540 slzdbm.wzm
23459 zhbwh.cds
$ cd ..
$ cd wrmr
$ ls
90357 zbdp.gqb
$ cd ..
$ cd zjrdgf
$ ls
301679 qbmt.vhh
12942 slfdb.bqt
$ cd ..
$ cd zrpqf
$ ls
241851 pnsdlbs.brw
$ cd ..
$ cd ..
$ cd nzcbqzdz
$ ls
182963 qgg.zsc
dir qwjgb
301344 slfdb.bqt
299029 zbdp.gqb
$ cd qwjgb
$ ls
223978 rprrjp.hfd
164106 wlq.rcq
$ cd ..
$ cd ..
$ cd ..
$ cd zdhcp
$ ls
dir hbggp
dir mzsfcvgv
dir zdsvth
$ cd hbggp
$ ls
302339 qcmfdhvm
112806 rpb.vfl
$ cd ..
$ cd mzsfcvgv
$ ls
dir gpz
dir slzdbm
dir zvjv
$ cd gpz
$ ls
290264 fdl
158548 slfdb.bqt
$ cd ..
$ cd slzdbm
$ ls
dir hlvpfw
dir mzsfcvgv
dir zdcthssc
$ cd hlvpfw
$ ls
49287 zbdp.gqb
$ cd ..
$ cd mzsfcvgv
$ ls
dir gnmtggs
103244 jnwr.gtn
263736 slfdb.bqt
$ cd gnmtggs
$ ls
31665 qhthp.lfh
$ cd ..
$ cd ..
$ cd zdcthssc
$ ls
64408 cgbnzc.cgn
$ cd ..
$ cd ..
$ cd zvjv
$ ls
6088 gnpsb.dml
145405 llp.cbm
dir mgcg
dir nwl
dir rqznjmt
128194 zbdp.gqb
$ cd mgcg
$ ls
dir ngwvm
dir slzdbm
dir smd
dir smr
dir zcgrgff
$ cd ngwvm
$ ls
49186 hmqfdht
$ cd ..
$ cd slzdbm
$ ls
44434 rdcb
$ cd ..
$ cd smd
$ ls
dir fjd
dir rscdvnwt
dir wbhp
318166 zbdp.gqb
$ cd fjd
$ ls
14084 hlvpfw.ghn
$ cd ..
$ cd rscdvnwt
$ ls
6432 fzwpmjm.ntl
$ cd ..
$ cd wbhp
$ ls
dir czv
217474 dnq.cpq
36009 jnwr.zdv
32123 mpq.cdn
27607 mrm.sgs
107593 zbdp.gqb
$ cd czv
$ ls
60750 zcgrgff.mnr
$ cd ..
$ cd ..
$ cd ..
$ cd smr
$ ls
47204 zcgrgff.vsl
$ cd ..
$ cd zcgrgff
$ ls
317050 hjtfb
110272 jrbcbvjw
290867 ltfj.dvl
dir mzsfcvgv
130694 wnc.rnr
206448 zbdp.gqb
$ cd mzsfcvgv
$ ls
dir jnwr
$ cd jnwr
$ ls
20405 bjpvssjt
$ cd ..
$ cd ..
$ cd ..
$ cd ..
$ cd nwl
$ ls
254439 dlv
193910 hlvpfw
13603 hnsdntn.trz
2990 mpq.cdn
$ cd ..
$ cd rqznjmt
$ ls
250114 hlvpfw
189732 mzsfcvgv.vqv
144688 zcgrgff
$ cd ..
$ cd ..
$ cd ..
$ cd zdsvth
$ ls
150609 jnwr.rtr
173402 nftm.nwd
dir njhjrgmf
$ cd njhjrgmf
$ ls
295762 nchztlh.lcs"""

