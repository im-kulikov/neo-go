// Code generated by "stringer -type=Opcode"; DO NOT EDIT.

package opcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[PUSH0-0]
	_ = x[PUSHF-0]
	_ = x[PUSHBYTES1-1]
	_ = x[PUSHBYTES2-2]
	_ = x[PUSHBYTES3-3]
	_ = x[PUSHBYTES4-4]
	_ = x[PUSHBYTES5-5]
	_ = x[PUSHBYTES6-6]
	_ = x[PUSHBYTES7-7]
	_ = x[PUSHBYTES8-8]
	_ = x[PUSHBYTES9-9]
	_ = x[PUSHBYTES10-10]
	_ = x[PUSHBYTES11-11]
	_ = x[PUSHBYTES12-12]
	_ = x[PUSHBYTES13-13]
	_ = x[PUSHBYTES14-14]
	_ = x[PUSHBYTES15-15]
	_ = x[PUSHBYTES16-16]
	_ = x[PUSHBYTES17-17]
	_ = x[PUSHBYTES18-18]
	_ = x[PUSHBYTES19-19]
	_ = x[PUSHBYTES20-20]
	_ = x[PUSHBYTES21-21]
	_ = x[PUSHBYTES22-22]
	_ = x[PUSHBYTES23-23]
	_ = x[PUSHBYTES24-24]
	_ = x[PUSHBYTES25-25]
	_ = x[PUSHBYTES26-26]
	_ = x[PUSHBYTES27-27]
	_ = x[PUSHBYTES28-28]
	_ = x[PUSHBYTES29-29]
	_ = x[PUSHBYTES30-30]
	_ = x[PUSHBYTES31-31]
	_ = x[PUSHBYTES32-32]
	_ = x[PUSHBYTES33-33]
	_ = x[PUSHBYTES34-34]
	_ = x[PUSHBYTES35-35]
	_ = x[PUSHBYTES36-36]
	_ = x[PUSHBYTES37-37]
	_ = x[PUSHBYTES38-38]
	_ = x[PUSHBYTES39-39]
	_ = x[PUSHBYTES40-40]
	_ = x[PUSHBYTES41-41]
	_ = x[PUSHBYTES42-42]
	_ = x[PUSHBYTES43-43]
	_ = x[PUSHBYTES44-44]
	_ = x[PUSHBYTES45-45]
	_ = x[PUSHBYTES46-46]
	_ = x[PUSHBYTES47-47]
	_ = x[PUSHBYTES48-48]
	_ = x[PUSHBYTES49-49]
	_ = x[PUSHBYTES50-50]
	_ = x[PUSHBYTES51-51]
	_ = x[PUSHBYTES52-52]
	_ = x[PUSHBYTES53-53]
	_ = x[PUSHBYTES54-54]
	_ = x[PUSHBYTES55-55]
	_ = x[PUSHBYTES56-56]
	_ = x[PUSHBYTES57-57]
	_ = x[PUSHBYTES58-58]
	_ = x[PUSHBYTES59-59]
	_ = x[PUSHBYTES60-60]
	_ = x[PUSHBYTES61-61]
	_ = x[PUSHBYTES62-62]
	_ = x[PUSHBYTES63-63]
	_ = x[PUSHBYTES64-64]
	_ = x[PUSHBYTES65-65]
	_ = x[PUSHBYTES66-66]
	_ = x[PUSHBYTES67-67]
	_ = x[PUSHBYTES68-68]
	_ = x[PUSHBYTES69-69]
	_ = x[PUSHBYTES70-70]
	_ = x[PUSHBYTES71-71]
	_ = x[PUSHBYTES72-72]
	_ = x[PUSHBYTES73-73]
	_ = x[PUSHBYTES74-74]
	_ = x[PUSHBYTES75-75]
	_ = x[PUSHDATA1-76]
	_ = x[PUSHDATA2-77]
	_ = x[PUSHDATA4-78]
	_ = x[PUSHM1-79]
	_ = x[PUSH1-81]
	_ = x[PUSHT-81]
	_ = x[PUSH2-82]
	_ = x[PUSH3-83]
	_ = x[PUSH4-84]
	_ = x[PUSH5-85]
	_ = x[PUSH6-86]
	_ = x[PUSH7-87]
	_ = x[PUSH8-88]
	_ = x[PUSH9-89]
	_ = x[PUSH10-90]
	_ = x[PUSH11-91]
	_ = x[PUSH12-92]
	_ = x[PUSH13-93]
	_ = x[PUSH14-94]
	_ = x[PUSH15-95]
	_ = x[PUSH16-96]
	_ = x[NOP-97]
	_ = x[JMP-98]
	_ = x[JMPIF-99]
	_ = x[JMPIFNOT-100]
	_ = x[CALL-101]
	_ = x[RET-102]
	_ = x[APPCALL-103]
	_ = x[SYSCALL-104]
	_ = x[TAILCALL-105]
	_ = x[DUPFROMALTSTACK-106]
	_ = x[TOALTSTACK-107]
	_ = x[FROMALTSTACK-108]
	_ = x[XDROP-109]
	_ = x[XSWAP-114]
	_ = x[XTUCK-115]
	_ = x[DEPTH-116]
	_ = x[DROP-117]
	_ = x[DUP-118]
	_ = x[NIP-119]
	_ = x[OVER-120]
	_ = x[PICK-121]
	_ = x[ROLL-122]
	_ = x[ROT-123]
	_ = x[SWAP-124]
	_ = x[TUCK-125]
	_ = x[CAT-126]
	_ = x[SUBSTR-127]
	_ = x[LEFT-128]
	_ = x[RIGHT-129]
	_ = x[SIZE-130]
	_ = x[INVERT-131]
	_ = x[AND-132]
	_ = x[OR-133]
	_ = x[XOR-134]
	_ = x[EQUAL-135]
	_ = x[INC-139]
	_ = x[DEC-140]
	_ = x[SIGN-141]
	_ = x[NEGATE-143]
	_ = x[ABS-144]
	_ = x[NOT-145]
	_ = x[NZ-146]
	_ = x[ADD-147]
	_ = x[SUB-148]
	_ = x[MUL-149]
	_ = x[DIV-150]
	_ = x[MOD-151]
	_ = x[SHL-152]
	_ = x[SHR-153]
	_ = x[BOOLAND-154]
	_ = x[BOOLOR-155]
	_ = x[NUMEQUAL-156]
	_ = x[NUMNOTEQUAL-158]
	_ = x[LT-159]
	_ = x[GT-160]
	_ = x[LTE-161]
	_ = x[GTE-162]
	_ = x[MIN-163]
	_ = x[MAX-164]
	_ = x[WITHIN-165]
	_ = x[SHA1-167]
	_ = x[SHA256-168]
	_ = x[HASH160-169]
	_ = x[HASH256-170]
	_ = x[CHECKSIG-172]
	_ = x[VERIFY-173]
	_ = x[CHECKMULTISIG-174]
	_ = x[ARRAYSIZE-192]
	_ = x[PACK-193]
	_ = x[UNPACK-194]
	_ = x[PICKITEM-195]
	_ = x[SETITEM-196]
	_ = x[NEWARRAY-197]
	_ = x[NEWSTRUCT-198]
	_ = x[NEWMAP-199]
	_ = x[APPEND-200]
	_ = x[REVERSE-201]
	_ = x[REMOVE-202]
	_ = x[HASKEY-203]
	_ = x[KEYS-204]
	_ = x[VALUES-205]
	_ = x[CALLI-224]
	_ = x[CALLE-225]
	_ = x[CALLED-226]
	_ = x[CALLET-227]
	_ = x[CALLEDT-228]
	_ = x[THROW-240]
	_ = x[THROWIFNOT-241]
}

const _Opcode_name = "PUSH0PUSHBYTES1PUSHBYTES2PUSHBYTES3PUSHBYTES4PUSHBYTES5PUSHBYTES6PUSHBYTES7PUSHBYTES8PUSHBYTES9PUSHBYTES10PUSHBYTES11PUSHBYTES12PUSHBYTES13PUSHBYTES14PUSHBYTES15PUSHBYTES16PUSHBYTES17PUSHBYTES18PUSHBYTES19PUSHBYTES20PUSHBYTES21PUSHBYTES22PUSHBYTES23PUSHBYTES24PUSHBYTES25PUSHBYTES26PUSHBYTES27PUSHBYTES28PUSHBYTES29PUSHBYTES30PUSHBYTES31PUSHBYTES32PUSHBYTES33PUSHBYTES34PUSHBYTES35PUSHBYTES36PUSHBYTES37PUSHBYTES38PUSHBYTES39PUSHBYTES40PUSHBYTES41PUSHBYTES42PUSHBYTES43PUSHBYTES44PUSHBYTES45PUSHBYTES46PUSHBYTES47PUSHBYTES48PUSHBYTES49PUSHBYTES50PUSHBYTES51PUSHBYTES52PUSHBYTES53PUSHBYTES54PUSHBYTES55PUSHBYTES56PUSHBYTES57PUSHBYTES58PUSHBYTES59PUSHBYTES60PUSHBYTES61PUSHBYTES62PUSHBYTES63PUSHBYTES64PUSHBYTES65PUSHBYTES66PUSHBYTES67PUSHBYTES68PUSHBYTES69PUSHBYTES70PUSHBYTES71PUSHBYTES72PUSHBYTES73PUSHBYTES74PUSHBYTES75PUSHDATA1PUSHDATA2PUSHDATA4PUSHM1PUSH1PUSH2PUSH3PUSH4PUSH5PUSH6PUSH7PUSH8PUSH9PUSH10PUSH11PUSH12PUSH13PUSH14PUSH15PUSH16NOPJMPJMPIFJMPIFNOTCALLRETAPPCALLSYSCALLTAILCALLDUPFROMALTSTACKTOALTSTACKFROMALTSTACKXDROPXSWAPXTUCKDEPTHDROPDUPNIPOVERPICKROLLROTSWAPTUCKCATSUBSTRLEFTRIGHTSIZEINVERTANDORXOREQUALINCDECSIGNNEGATEABSNOTNZADDSUBMULDIVMODSHLSHRBOOLANDBOOLORNUMEQUALNUMNOTEQUALLTGTLTEGTEMINMAXWITHINSHA1SHA256HASH160HASH256CHECKSIGVERIFYCHECKMULTISIGARRAYSIZEPACKUNPACKPICKITEMSETITEMNEWARRAYNEWSTRUCTNEWMAPAPPENDREVERSEREMOVEHASKEYKEYSVALUESCALLICALLECALLEDCALLETCALLEDTTHROWTHROWIFNOT"

var _Opcode_map = map[Opcode]string{
	0:   _Opcode_name[0:5],
	1:   _Opcode_name[5:15],
	2:   _Opcode_name[15:25],
	3:   _Opcode_name[25:35],
	4:   _Opcode_name[35:45],
	5:   _Opcode_name[45:55],
	6:   _Opcode_name[55:65],
	7:   _Opcode_name[65:75],
	8:   _Opcode_name[75:85],
	9:   _Opcode_name[85:95],
	10:  _Opcode_name[95:106],
	11:  _Opcode_name[106:117],
	12:  _Opcode_name[117:128],
	13:  _Opcode_name[128:139],
	14:  _Opcode_name[139:150],
	15:  _Opcode_name[150:161],
	16:  _Opcode_name[161:172],
	17:  _Opcode_name[172:183],
	18:  _Opcode_name[183:194],
	19:  _Opcode_name[194:205],
	20:  _Opcode_name[205:216],
	21:  _Opcode_name[216:227],
	22:  _Opcode_name[227:238],
	23:  _Opcode_name[238:249],
	24:  _Opcode_name[249:260],
	25:  _Opcode_name[260:271],
	26:  _Opcode_name[271:282],
	27:  _Opcode_name[282:293],
	28:  _Opcode_name[293:304],
	29:  _Opcode_name[304:315],
	30:  _Opcode_name[315:326],
	31:  _Opcode_name[326:337],
	32:  _Opcode_name[337:348],
	33:  _Opcode_name[348:359],
	34:  _Opcode_name[359:370],
	35:  _Opcode_name[370:381],
	36:  _Opcode_name[381:392],
	37:  _Opcode_name[392:403],
	38:  _Opcode_name[403:414],
	39:  _Opcode_name[414:425],
	40:  _Opcode_name[425:436],
	41:  _Opcode_name[436:447],
	42:  _Opcode_name[447:458],
	43:  _Opcode_name[458:469],
	44:  _Opcode_name[469:480],
	45:  _Opcode_name[480:491],
	46:  _Opcode_name[491:502],
	47:  _Opcode_name[502:513],
	48:  _Opcode_name[513:524],
	49:  _Opcode_name[524:535],
	50:  _Opcode_name[535:546],
	51:  _Opcode_name[546:557],
	52:  _Opcode_name[557:568],
	53:  _Opcode_name[568:579],
	54:  _Opcode_name[579:590],
	55:  _Opcode_name[590:601],
	56:  _Opcode_name[601:612],
	57:  _Opcode_name[612:623],
	58:  _Opcode_name[623:634],
	59:  _Opcode_name[634:645],
	60:  _Opcode_name[645:656],
	61:  _Opcode_name[656:667],
	62:  _Opcode_name[667:678],
	63:  _Opcode_name[678:689],
	64:  _Opcode_name[689:700],
	65:  _Opcode_name[700:711],
	66:  _Opcode_name[711:722],
	67:  _Opcode_name[722:733],
	68:  _Opcode_name[733:744],
	69:  _Opcode_name[744:755],
	70:  _Opcode_name[755:766],
	71:  _Opcode_name[766:777],
	72:  _Opcode_name[777:788],
	73:  _Opcode_name[788:799],
	74:  _Opcode_name[799:810],
	75:  _Opcode_name[810:821],
	76:  _Opcode_name[821:830],
	77:  _Opcode_name[830:839],
	78:  _Opcode_name[839:848],
	79:  _Opcode_name[848:854],
	81:  _Opcode_name[854:859],
	82:  _Opcode_name[859:864],
	83:  _Opcode_name[864:869],
	84:  _Opcode_name[869:874],
	85:  _Opcode_name[874:879],
	86:  _Opcode_name[879:884],
	87:  _Opcode_name[884:889],
	88:  _Opcode_name[889:894],
	89:  _Opcode_name[894:899],
	90:  _Opcode_name[899:905],
	91:  _Opcode_name[905:911],
	92:  _Opcode_name[911:917],
	93:  _Opcode_name[917:923],
	94:  _Opcode_name[923:929],
	95:  _Opcode_name[929:935],
	96:  _Opcode_name[935:941],
	97:  _Opcode_name[941:944],
	98:  _Opcode_name[944:947],
	99:  _Opcode_name[947:952],
	100: _Opcode_name[952:960],
	101: _Opcode_name[960:964],
	102: _Opcode_name[964:967],
	103: _Opcode_name[967:974],
	104: _Opcode_name[974:981],
	105: _Opcode_name[981:989],
	106: _Opcode_name[989:1004],
	107: _Opcode_name[1004:1014],
	108: _Opcode_name[1014:1026],
	109: _Opcode_name[1026:1031],
	114: _Opcode_name[1031:1036],
	115: _Opcode_name[1036:1041],
	116: _Opcode_name[1041:1046],
	117: _Opcode_name[1046:1050],
	118: _Opcode_name[1050:1053],
	119: _Opcode_name[1053:1056],
	120: _Opcode_name[1056:1060],
	121: _Opcode_name[1060:1064],
	122: _Opcode_name[1064:1068],
	123: _Opcode_name[1068:1071],
	124: _Opcode_name[1071:1075],
	125: _Opcode_name[1075:1079],
	126: _Opcode_name[1079:1082],
	127: _Opcode_name[1082:1088],
	128: _Opcode_name[1088:1092],
	129: _Opcode_name[1092:1097],
	130: _Opcode_name[1097:1101],
	131: _Opcode_name[1101:1107],
	132: _Opcode_name[1107:1110],
	133: _Opcode_name[1110:1112],
	134: _Opcode_name[1112:1115],
	135: _Opcode_name[1115:1120],
	139: _Opcode_name[1120:1123],
	140: _Opcode_name[1123:1126],
	141: _Opcode_name[1126:1130],
	143: _Opcode_name[1130:1136],
	144: _Opcode_name[1136:1139],
	145: _Opcode_name[1139:1142],
	146: _Opcode_name[1142:1144],
	147: _Opcode_name[1144:1147],
	148: _Opcode_name[1147:1150],
	149: _Opcode_name[1150:1153],
	150: _Opcode_name[1153:1156],
	151: _Opcode_name[1156:1159],
	152: _Opcode_name[1159:1162],
	153: _Opcode_name[1162:1165],
	154: _Opcode_name[1165:1172],
	155: _Opcode_name[1172:1178],
	156: _Opcode_name[1178:1186],
	158: _Opcode_name[1186:1197],
	159: _Opcode_name[1197:1199],
	160: _Opcode_name[1199:1201],
	161: _Opcode_name[1201:1204],
	162: _Opcode_name[1204:1207],
	163: _Opcode_name[1207:1210],
	164: _Opcode_name[1210:1213],
	165: _Opcode_name[1213:1219],
	167: _Opcode_name[1219:1223],
	168: _Opcode_name[1223:1229],
	169: _Opcode_name[1229:1236],
	170: _Opcode_name[1236:1243],
	172: _Opcode_name[1243:1251],
	173: _Opcode_name[1251:1257],
	174: _Opcode_name[1257:1270],
	192: _Opcode_name[1270:1279],
	193: _Opcode_name[1279:1283],
	194: _Opcode_name[1283:1289],
	195: _Opcode_name[1289:1297],
	196: _Opcode_name[1297:1304],
	197: _Opcode_name[1304:1312],
	198: _Opcode_name[1312:1321],
	199: _Opcode_name[1321:1327],
	200: _Opcode_name[1327:1333],
	201: _Opcode_name[1333:1340],
	202: _Opcode_name[1340:1346],
	203: _Opcode_name[1346:1352],
	204: _Opcode_name[1352:1356],
	205: _Opcode_name[1356:1362],
	224: _Opcode_name[1362:1367],
	225: _Opcode_name[1367:1372],
	226: _Opcode_name[1372:1378],
	227: _Opcode_name[1378:1384],
	228: _Opcode_name[1384:1391],
	240: _Opcode_name[1391:1396],
	241: _Opcode_name[1396:1406],
}

func (i Opcode) String() string {
	if str, ok := _Opcode_map[i]; ok {
		return str
	}
	return "Opcode(" + strconv.FormatInt(int64(i), 10) + ")"
}
