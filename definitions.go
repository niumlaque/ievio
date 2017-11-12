package ievio

import (
	"fmt"
	"strconv"
	"strings"
)

type EV_TYPE uint16

const (
	EV_SYN       EV_TYPE = 0x00
	EV_KEY               = 0x01
	EV_REL               = 0x02
	EV_ABS               = 0x03
	EV_MSC               = 0x04
	EV_SW                = 0x05
	EV_LED               = 0x11
	EV_SND               = 0x12
	EV_REP               = 0x14
	EV_FF                = 0x15
	EV_PWR               = 0x16
	EV_FF_STATUS         = 0x17
	EV_MAX               = 0x1f
	EV_CNT               = (EV_MAX + 1)
)

func stringToUint64(s string, bitSize int) (uint64, error) {
	if strings.HasPrefix(strings.ToLower(s), "0x") {
		return strconv.ParseUint(string([]rune(s)[2:]), 16, bitSize)
	} else {
		return strconv.ParseUint(s, 10, bitSize)
	}
}

func (v EV_TYPE) String() string {
	switch v {
	case EV_SYN:
		return fmt.Sprintf("EV_SYN(0x%04x)", uint16(v))
	case EV_KEY:
		return fmt.Sprintf("EV_KEY(0x%04x)", uint16(v))
	case EV_REL:
		return fmt.Sprintf("EV_REL(0x%04x)", uint16(v))
	case EV_ABS:
		return fmt.Sprintf("EV_ABS(0x%04x)", uint16(v))
	case EV_MSC:
		return fmt.Sprintf("EV_MSC(0x%04x)", uint16(v))
	case EV_SW:
		return fmt.Sprintf("EV_SW(0x%04x)", uint16(v))
	case EV_LED:
		return fmt.Sprintf("EV_LED(0x%04x)", uint16(v))
	case EV_SND:
		return fmt.Sprintf("EV_SND(0x%04x)", uint16(v))
	case EV_REP:
		return fmt.Sprintf("EV_REP(0x%04x)", uint16(v))
	case EV_FF:
		return fmt.Sprintf("EV_FF(0x%04x)", uint16(v))
	case EV_PWR:
		return fmt.Sprintf("EV_PWR(0x%04x)", uint16(v))
	case EV_FF_STATUS:
		return fmt.Sprintf("EV_FF_STATUS(0x%04x)", uint16(v))
	case EV_MAX:
		return fmt.Sprintf("EV_MAX(0x%04x)", uint16(v))
	case EV_CNT:
		return fmt.Sprintf("EV_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

func getEventTypeFromString(s string) (EV_TYPE, error) {
	v, err := stringToUint64(s, 16)
	if err != nil {
		return EV_CNT, err
	}

	return EV_TYPE(uint16(v)), nil
}

type Code interface {
	ValueUint16() uint16
	String() string
}

func getCodeFromString(eventType EV_TYPE, s string) (Code, error) {
	v, err := stringToUint64(s, 16)
	if err != nil {
		return nil, err
	}

	switch eventType {
	case EV_SYN:
		return NewSynCode(SYN_CODE(v)), nil
	case EV_KEY:
		return NewKeyCode(KEY_CODE(v)), nil
	case EV_REL:
		return NewRelCode(REL_CODE(v)), nil
	case EV_ABS:
		return NewAbsCode(ABS_CODE(v)), nil
	case EV_MSC:
		return NewMscCode(MSC_CODE(v)), nil
	case EV_SW:
		return NewSwCode(SW_CODE(v)), nil
	case EV_LED:
		return NewLedCode(LED_CODE(v)), nil
	case EV_SND:
		return NewSndCode(SND_CODE(v)), nil
	case EV_REP:
		return NewRepCode(REP_CODE(v)), nil
	case EV_FF:
		return nil, nil // FIXME
	case EV_PWR:
		return nil, nil // FIXME
	case EV_FF_STATUS:
		return nil, nil // FIXME
	case EV_MAX:
		return nil, nil // FIXME
	case EV_CNT:
		return nil, nil // FIXME
	}

	return nil, nil
}

type SYN_CODE uint16
type SynCode struct {
	v SYN_CODE
}

func NewSynCode(v SYN_CODE) *SynCode {
	return &SynCode{
		v: v,
	}
}

func (v *SynCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *SynCode) String() string {
	return v.v.String()
}

const (
	SYN_REPORT    SYN_CODE = 0
	SYN_CONFIG             = 1
	SYN_MT_REPORT          = 2
	SYN_DROPPED            = 3
	SYN_MAX                = 0xf
	SYN_CNT                = (SYN_MAX + 1)
)

func (v SYN_CODE) String() string {
	switch v {
	case SYN_REPORT:
		return fmt.Sprintf("SYN_REPORT(0x%04x)", uint16(v))
	case SYN_CONFIG:
		return fmt.Sprintf("SYN_CONFIG(0x%04x)", uint16(v))
	case SYN_MT_REPORT:
		return fmt.Sprintf("SYN_MT_REPORT(0x%04x)", uint16(v))
	case SYN_DROPPED:
		return fmt.Sprintf("SYN_DROPPED(0x%04x)", uint16(v))
	case SYN_MAX:
		return fmt.Sprintf("SYN_MAX(0x%04x)", uint16(v))
	case SYN_CNT:
		return fmt.Sprintf("SYN_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type KEY_CODE uint16
type KeyCode struct {
	v KEY_CODE
}

func NewKeyCode(v KEY_CODE) *KeyCode {
	return &KeyCode{
		v: v,
	}
}

func (v *KeyCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *KeyCode) String() string {
	return v.v.String()
}

const (
	KEY_RESERVED                 KEY_CODE = 0
	KEY_ESC                               = 1
	KEY_1                                 = 2
	KEY_2                                 = 3
	KEY_3                                 = 4
	KEY_4                                 = 5
	KEY_5                                 = 6
	KEY_6                                 = 7
	KEY_7                                 = 8
	KEY_8                                 = 9
	KEY_9                                 = 10
	KEY_0                                 = 11
	KEY_MINUS                             = 12
	KEY_EQUAL                             = 13
	KEY_BACKSPACE                         = 14
	KEY_TAB                               = 15
	KEY_Q                                 = 16
	KEY_W                                 = 17
	KEY_E                                 = 18
	KEY_R                                 = 19
	KEY_T                                 = 20
	KEY_Y                                 = 21
	KEY_U                                 = 22
	KEY_I                                 = 23
	KEY_O                                 = 24
	KEY_P                                 = 25
	KEY_LEFTBRACE                         = 26
	KEY_RIGHTBRACE                        = 27
	KEY_ENTER                             = 28
	KEY_LEFTCTRL                          = 29
	KEY_A                                 = 30
	KEY_S                                 = 31
	KEY_D                                 = 32
	KEY_F                                 = 33
	KEY_G                                 = 34
	KEY_H                                 = 35
	KEY_J                                 = 36
	KEY_K                                 = 37
	KEY_L                                 = 38
	KEY_SEMICOLON                         = 39
	KEY_APOSTROPHE                        = 40
	KEY_GRAVE                             = 41
	KEY_LEFTSHIFT                         = 42
	KEY_BACKSLASH                         = 43
	KEY_Z                                 = 44
	KEY_X                                 = 45
	KEY_C                                 = 46
	KEY_V                                 = 47
	KEY_B                                 = 48
	KEY_N                                 = 49
	KEY_M                                 = 50
	KEY_COMMA                             = 51
	KEY_DOT                               = 52
	KEY_SLASH                             = 53
	KEY_RIGHTSHIFT                        = 54
	KEY_KPASTERISK                        = 55
	KEY_LEFTALT                           = 56
	KEY_SPACE                             = 57
	KEY_CAPSLOCK                          = 58
	KEY_F1                                = 59
	KEY_F2                                = 60
	KEY_F3                                = 61
	KEY_F4                                = 62
	KEY_F5                                = 63
	KEY_F6                                = 64
	KEY_F7                                = 65
	KEY_F8                                = 66
	KEY_F9                                = 67
	KEY_F10                               = 68
	KEY_NUMLOCK                           = 69
	KEY_SCROLLLOCK                        = 70
	KEY_KP7                               = 71
	KEY_KP8                               = 72
	KEY_KP9                               = 73
	KEY_KPMINUS                           = 74
	KEY_KP4                               = 75
	KEY_KP5                               = 76
	KEY_KP6                               = 77
	KEY_KPPLUS                            = 78
	KEY_KP1                               = 79
	KEY_KP2                               = 80
	KEY_KP3                               = 81
	KEY_KP0                               = 82
	KEY_KPDOT                             = 83
	KEY_ZENKAKUHANKAKU                    = 85
	KEY_102ND                             = 86
	KEY_F11                               = 87
	KEY_F12                               = 88
	KEY_RO                                = 89
	KEY_KATAKANA                          = 90
	KEY_HIRAGANA                          = 91
	KEY_HENKAN                            = 92
	KEY_KATAKANAHIRAGANA                  = 93
	KEY_MUHENKAN                          = 94
	KEY_KPJPCOMMA                         = 95
	KEY_KPENTER                           = 96
	KEY_RIGHTCTRL                         = 97
	KEY_KPSLASH                           = 98
	KEY_SYSRQ                             = 99
	KEY_RIGHTALT                          = 100
	KEY_LINEFEED                          = 101
	KEY_HOME                              = 102
	KEY_UP                                = 103
	KEY_PAGEUP                            = 104
	KEY_LEFT                              = 105
	KEY_RIGHT                             = 106
	KEY_END                               = 107
	KEY_DOWN                              = 108
	KEY_PAGEDOWN                          = 109
	KEY_INSERT                            = 110
	KEY_DELETE                            = 111
	KEY_MACRO                             = 112
	KEY_MUTE                              = 113
	KEY_VOLUMEDOWN                        = 114
	KEY_VOLUMEUP                          = 115
	KEY_POWER                             = 116
	KEY_KPEQUAL                           = 117
	KEY_KPPLUSMINUS                       = 118
	KEY_PAUSE                             = 119
	KEY_SCALE                             = 120
	KEY_KPCOMMA                           = 121
	KEY_HANGEUL                           = 122
	KEY_HANGUEL                           = KEY_HANGEUL
	KEY_HANJA                             = 123
	KEY_YEN                               = 124
	KEY_LEFTMETA                          = 125
	KEY_RIGHTMETA                         = 126
	KEY_COMPOSE                           = 127
	KEY_STOP                              = 128
	KEY_AGAIN                             = 129
	KEY_PROPS                             = 130
	KEY_UNDO                              = 131
	KEY_FRONT                             = 132
	KEY_COPY                              = 133
	KEY_OPEN                              = 134
	KEY_PASTE                             = 135
	KEY_FIND                              = 136
	KEY_CUT                               = 137
	KEY_HELP                              = 138
	KEY_MENU                              = 139
	KEY_CALC                              = 140
	KEY_SETUP                             = 141
	KEY_SLEEP                             = 142
	KEY_WAKEUP                            = 143
	KEY_FILE                              = 144
	KEY_SENDFILE                          = 145
	KEY_DELETEFILE                        = 146
	KEY_XFER                              = 147
	KEY_PROG1                             = 148
	KEY_PROG2                             = 149
	KEY_WWW                               = 150
	KEY_MSDOS                             = 151
	KEY_COFFEE                            = 152
	KEY_SCREENLOCK                        = KEY_COFFEE
	KEY_ROTATE_DISPLAY                    = 153
	KEY_DIRECTION                         = KEY_ROTATE_DISPLAY
	KEY_CYCLEWINDOWS                      = 154
	KEY_MAIL                              = 155
	KEY_BOOKMARKS                         = 156
	KEY_COMPUTER                          = 157
	KEY_BACK                              = 158
	KEY_FORWARD                           = 159
	KEY_CLOSECD                           = 160
	KEY_EJECTCD                           = 161
	KEY_EJECTCLOSECD                      = 162
	KEY_NEXTSONG                          = 163
	KEY_PLAYPAUSE                         = 164
	KEY_PREVIOUSSONG                      = 165
	KEY_STOPCD                            = 166
	KEY_RECORD                            = 167
	KEY_REWIND                            = 168
	KEY_PHONE                             = 169
	KEY_ISO                               = 170
	KEY_CONFIG                            = 171
	KEY_HOMEPAGE                          = 172
	KEY_REFRESH                           = 173
	KEY_EXIT                              = 174
	KEY_MOVE                              = 175
	KEY_EDIT                              = 176
	KEY_SCROLLUP                          = 177
	KEY_SCROLLDOWN                        = 178
	KEY_KPLEFTPAREN                       = 179
	KEY_KPRIGHTPAREN                      = 180
	KEY_NEW                               = 181
	KEY_REDO                              = 182
	KEY_F13                               = 183
	KEY_F14                               = 184
	KEY_F15                               = 185
	KEY_F16                               = 186
	KEY_F17                               = 187
	KEY_F18                               = 188
	KEY_F19                               = 189
	KEY_F20                               = 190
	KEY_F21                               = 191
	KEY_F22                               = 192
	KEY_F23                               = 193
	KEY_F24                               = 194
	KEY_PLAYCD                            = 200
	KEY_PAUSECD                           = 201
	KEY_PROG3                             = 202
	KEY_PROG4                             = 203
	KEY_DASHBOARD                         = 204
	KEY_SUSPEND                           = 205
	KEY_CLOSE                             = 206
	KEY_PLAY                              = 207
	KEY_FASTFORWARD                       = 208
	KEY_BASSBOOST                         = 209
	KEY_PRINT                             = 210
	KEY_HP                                = 211
	KEY_CAMERA                            = 212
	KEY_SOUND                             = 213
	KEY_QUESTION                          = 214
	KEY_EMAIL                             = 215
	KEY_CHAT                              = 216
	KEY_SEARCH                            = 217
	KEY_CONNECT                           = 218
	KEY_FINANCE                           = 219
	KEY_SPORT                             = 220
	KEY_SHOP                              = 221
	KEY_ALTERASE                          = 222
	KEY_CANCEL                            = 223
	KEY_BRIGHTNESSDOWN                    = 224
	KEY_BRIGHTNESSUP                      = 225
	KEY_MEDIA                             = 226
	KEY_SWITCHVIDEOMODE                   = 227
	KEY_KBDILLUMTOGGLE                    = 228
	KEY_KBDILLUMDOWN                      = 229
	KEY_KBDILLUMUP                        = 230
	KEY_SEND                              = 231
	KEY_REPLY                             = 232
	KEY_FORWARDMAIL                       = 233
	KEY_SAVE                              = 234
	KEY_DOCUMENTS                         = 235
	KEY_BATTERY                           = 236
	KEY_BLUETOOTH                         = 237
	KEY_WLAN                              = 238
	KEY_UWB                               = 239
	KEY_UNKNOWN                           = 240
	KEY_VIDEO_NEXT                        = 241
	KEY_VIDEO_PREV                        = 242
	KEY_BRIGHTNESS_CYCLE                  = 243
	KEY_BRIGHTNESS_AUTO                   = 244
	KEY_BRIGHTNESS_ZERO                   = KEY_BRIGHTNESS_AUTO
	KEY_DISPLAY_OFF                       = 245
	KEY_WWAN                              = 246
	KEY_WIMAX                             = KEY_WWAN
	KEY_RFKILL                            = 247
	KEY_MICMUTE                           = 248
	BTN_MISC                              = 0x100
	BTN_0                                 = 0x100
	BTN_1                                 = 0x101
	BTN_2                                 = 0x102
	BTN_3                                 = 0x103
	BTN_4                                 = 0x104
	BTN_5                                 = 0x105
	BTN_6                                 = 0x106
	BTN_7                                 = 0x107
	BTN_8                                 = 0x108
	BTN_9                                 = 0x109
	BTN_MOUSE                             = 0x110
	BTN_LEFT                              = 0x110
	BTN_RIGHT                             = 0x111
	BTN_MIDDLE                            = 0x112
	BTN_SIDE                              = 0x113
	BTN_EXTRA                             = 0x114
	BTN_FORWARD                           = 0x115
	BTN_BACK                              = 0x116
	BTN_TASK                              = 0x117
	BTN_JOYSTICK                          = 0x120
	BTN_TRIGGER                           = 0x120
	BTN_THUMB                             = 0x121
	BTN_THUMB2                            = 0x122
	BTN_TOP                               = 0x123
	BTN_TOP2                              = 0x124
	BTN_PINKIE                            = 0x125
	BTN_BASE                              = 0x126
	BTN_BASE2                             = 0x127
	BTN_BASE3                             = 0x128
	BTN_BASE4                             = 0x129
	BTN_BASE5                             = 0x12a
	BTN_BASE6                             = 0x12b
	BTN_DEAD                              = 0x12f
	BTN_GAMEPAD                           = 0x130
	BTN_SOUTH                             = 0x130
	BTN_A                                 = BTN_SOUTH
	BTN_EAST                              = 0x131
	BTN_B                                 = BTN_EAST
	BTN_C                                 = 0x132
	BTN_NORTH                             = 0x133
	BTN_X                                 = BTN_NORTH
	BTN_WEST                              = 0x134
	BTN_Y                                 = BTN_WEST
	BTN_Z                                 = 0x135
	BTN_TL                                = 0x136
	BTN_TR                                = 0x137
	BTN_TL2                               = 0x138
	BTN_TR2                               = 0x139
	BTN_SELECT                            = 0x13a
	BTN_START                             = 0x13b
	BTN_MODE                              = 0x13c
	BTN_THUMBL                            = 0x13d
	BTN_THUMBR                            = 0x13e
	BTN_DIGI                              = 0x140
	BTN_TOOL_PEN                          = 0x140
	BTN_TOOL_RUBBER                       = 0x141
	BTN_TOOL_BRUSH                        = 0x142
	BTN_TOOL_PENCIL                       = 0x143
	BTN_TOOL_AIRBRUSH                     = 0x144
	BTN_TOOL_FINGER                       = 0x145
	BTN_TOOL_MOUSE                        = 0x146
	BTN_TOOL_LENS                         = 0x147
	BTN_TOOL_QUINTTAP                     = 0x148
	BTN_TOUCH                             = 0x14a
	BTN_STYLUS                            = 0x14b
	BTN_STYLUS2                           = 0x14c
	BTN_TOOL_DOUBLETAP                    = 0x14d
	BTN_TOOL_TRIPLETAP                    = 0x14e
	BTN_TOOL_QUADTAP                      = 0x14f
	BTN_WHEEL                             = 0x150
	BTN_GEAR_DOWN                         = 0x150
	BTN_GEAR_UP                           = 0x151
	KEY_OK                                = 0x160
	KEY_SELECT                            = 0x161
	KEY_GOTO                              = 0x162
	KEY_CLEAR                             = 0x163
	KEY_POWER2                            = 0x164
	KEY_OPTION                            = 0x165
	KEY_INFO                              = 0x166
	KEY_TIME                              = 0x167
	KEY_VENDOR                            = 0x168
	KEY_ARCHIVE                           = 0x169
	KEY_PROGRAM                           = 0x16a
	KEY_CHANNEL                           = 0x16b
	KEY_FAVORITES                         = 0x16c
	KEY_EPG                               = 0x16d
	KEY_PVR                               = 0x16e
	KEY_MHP                               = 0x16f
	KEY_LANGUAGE                          = 0x170
	KEY_TITLE                             = 0x171
	KEY_SUBTITLE                          = 0x172
	KEY_ANGLE                             = 0x173
	KEY_ZOOM                              = 0x174
	KEY_MODE                              = 0x175
	KEY_KEYBOARD                          = 0x176
	KEY_SCREEN                            = 0x177
	KEY_PC                                = 0x178
	KEY_TV                                = 0x179
	KEY_TV2                               = 0x17a
	KEY_VCR                               = 0x17b
	KEY_VCR2                              = 0x17c
	KEY_SAT                               = 0x17d
	KEY_SAT2                              = 0x17e
	KEY_CD                                = 0x17f
	KEY_TAPE                              = 0x180
	KEY_RADIO                             = 0x181
	KEY_TUNER                             = 0x182
	KEY_PLAYER                            = 0x183
	KEY_TEXT                              = 0x184
	KEY_DVD                               = 0x185
	KEY_AUX                               = 0x186
	KEY_MP3                               = 0x187
	KEY_AUDIO                             = 0x188
	KEY_VIDEO                             = 0x189
	KEY_DIRECTORY                         = 0x18a
	KEY_LIST                              = 0x18b
	KEY_MEMO                              = 0x18c
	KEY_CALENDAR                          = 0x18d
	KEY_RED                               = 0x18e
	KEY_GREEN                             = 0x18f
	KEY_YELLOW                            = 0x190
	KEY_BLUE                              = 0x191
	KEY_CHANNELUP                         = 0x192
	KEY_CHANNELDOWN                       = 0x193
	KEY_FIRST                             = 0x194
	KEY_LAST                              = 0x195
	KEY_AB                                = 0x196
	KEY_NEXT                              = 0x197
	KEY_RESTART                           = 0x198
	KEY_SLOW                              = 0x199
	KEY_SHUFFLE                           = 0x19a
	KEY_BREAK                             = 0x19b
	KEY_PREVIOUS                          = 0x19c
	KEY_DIGITS                            = 0x19d
	KEY_TEEN                              = 0x19e
	KEY_TWEN                              = 0x19f
	KEY_VIDEOPHONE                        = 0x1a0
	KEY_GAMES                             = 0x1a1
	KEY_ZOOMIN                            = 0x1a2
	KEY_ZOOMOUT                           = 0x1a3
	KEY_ZOOMRESET                         = 0x1a4
	KEY_WORDPROCESSOR                     = 0x1a5
	KEY_EDITOR                            = 0x1a6
	KEY_SPREADSHEET                       = 0x1a7
	KEY_GRAPHICSEDITOR                    = 0x1a8
	KEY_PRESENTATION                      = 0x1a9
	KEY_DATABASE                          = 0x1aa
	KEY_NEWS                              = 0x1ab
	KEY_VOICEMAIL                         = 0x1ac
	KEY_ADDRESSBOOK                       = 0x1ad
	KEY_MESSENGER                         = 0x1ae
	KEY_DISPLAYTOGGLE                     = 0x1af
	KEY_BRIGHTNESS_TOGGLE                 = KEY_DISPLAYTOGGLE
	KEY_SPELLCHECK                        = 0x1b0
	KEY_LOGOFF                            = 0x1b1
	KEY_DOLLAR                            = 0x1b2
	KEY_EURO                              = 0x1b3
	KEY_FRAMEBACK                         = 0x1b4
	KEY_FRAMEFORWARD                      = 0x1b5
	KEY_CONTEXT_MENU                      = 0x1b6
	KEY_MEDIA_REPEAT                      = 0x1b7
	KEY_10CHANNELSUP                      = 0x1b8
	KEY_10CHANNELSDOWN                    = 0x1b9
	KEY_IMAGES                            = 0x1ba
	KEY_DEL_EOL                           = 0x1c0
	KEY_DEL_EOS                           = 0x1c1
	KEY_INS_LINE                          = 0x1c2
	KEY_DEL_LINE                          = 0x1c3
	KEY_FN                                = 0x1d0
	KEY_FN_ESC                            = 0x1d1
	KEY_FN_F1                             = 0x1d2
	KEY_FN_F2                             = 0x1d3
	KEY_FN_F3                             = 0x1d4
	KEY_FN_F4                             = 0x1d5
	KEY_FN_F5                             = 0x1d6
	KEY_FN_F6                             = 0x1d7
	KEY_FN_F7                             = 0x1d8
	KEY_FN_F8                             = 0x1d9
	KEY_FN_F9                             = 0x1da
	KEY_FN_F10                            = 0x1db
	KEY_FN_F11                            = 0x1dc
	KEY_FN_F12                            = 0x1dd
	KEY_FN_1                              = 0x1de
	KEY_FN_2                              = 0x1df
	KEY_FN_D                              = 0x1e0
	KEY_FN_E                              = 0x1e1
	KEY_FN_F                              = 0x1e2
	KEY_FN_S                              = 0x1e3
	KEY_FN_B                              = 0x1e4
	KEY_BRL_DOT1                          = 0x1f1
	KEY_BRL_DOT2                          = 0x1f2
	KEY_BRL_DOT3                          = 0x1f3
	KEY_BRL_DOT4                          = 0x1f4
	KEY_BRL_DOT5                          = 0x1f5
	KEY_BRL_DOT6                          = 0x1f6
	KEY_BRL_DOT7                          = 0x1f7
	KEY_BRL_DOT8                          = 0x1f8
	KEY_BRL_DOT9                          = 0x1f9
	KEY_BRL_DOT10                         = 0x1fa
	KEY_NUMERIC_0                         = 0x200
	KEY_NUMERIC_1                         = 0x201
	KEY_NUMERIC_2                         = 0x202
	KEY_NUMERIC_3                         = 0x203
	KEY_NUMERIC_4                         = 0x204
	KEY_NUMERIC_5                         = 0x205
	KEY_NUMERIC_6                         = 0x206
	KEY_NUMERIC_7                         = 0x207
	KEY_NUMERIC_8                         = 0x208
	KEY_NUMERIC_9                         = 0x209
	KEY_NUMERIC_STAR                      = 0x20a
	KEY_NUMERIC_POUND                     = 0x20b
	KEY_NUMERIC_A                         = 0x20c
	KEY_NUMERIC_B                         = 0x20d
	KEY_NUMERIC_C                         = 0x20e
	KEY_NUMERIC_D                         = 0x20f
	KEY_CAMERA_FOCUS                      = 0x210
	KEY_WPS_BUTTON                        = 0x211
	KEY_TOUCHPAD_TOGGLE                   = 0x212
	KEY_TOUCHPAD_ON                       = 0x213
	KEY_TOUCHPAD_OFF                      = 0x214
	KEY_CAMERA_ZOOMIN                     = 0x215
	KEY_CAMERA_ZOOMOUT                    = 0x216
	KEY_CAMERA_UP                         = 0x217
	KEY_CAMERA_DOWN                       = 0x218
	KEY_CAMERA_LEFT                       = 0x219
	KEY_CAMERA_RIGHT                      = 0x21a
	KEY_ATTENDANT_ON                      = 0x21b
	KEY_ATTENDANT_OFF                     = 0x21c
	KEY_ATTENDANT_TOGGLE                  = 0x21d
	KEY_LIGHTS_TOGGLE                     = 0x21e
	BTN_DPAD_UP                           = 0x220
	BTN_DPAD_DOWN                         = 0x221
	BTN_DPAD_LEFT                         = 0x222
	BTN_DPAD_RIGHT                        = 0x223
	KEY_ALS_TOGGLE                        = 0x230
	KEY_BUTTONCONFIG                      = 0x240
	KEY_TASKMANAGER                       = 0x241
	KEY_JOURNAL                           = 0x242
	KEY_CONTROLPANEL                      = 0x243
	KEY_APPSELECT                         = 0x244
	KEY_SCREENSAVER                       = 0x245
	KEY_VOICECOMMAND                      = 0x246
	KEY_ASSISTANT                         = 0x247
	KEY_BRIGHTNESS_MIN                    = 0x250
	KEY_BRIGHTNESS_MAX                    = 0x251
	KEY_KBDINPUTASSIST_PREV               = 0x260
	KEY_KBDINPUTASSIST_NEXT               = 0x261
	KEY_KBDINPUTASSIST_PREVGROUP          = 0x262
	KEY_KBDINPUTASSIST_NEXTGROUP          = 0x263
	KEY_KBDINPUTASSIST_ACCEPT             = 0x264
	KEY_KBDINPUTASSIST_CANCEL             = 0x265
	KEY_RIGHT_UP                          = 0x266
	KEY_RIGHT_DOWN                        = 0x267
	KEY_LEFT_UP                           = 0x268
	KEY_LEFT_DOWN                         = 0x269
	KEY_ROOT_MENU                         = 0x26a
	KEY_MEDIA_TOP_MENU                    = 0x26b
	KEY_NUMERIC_11                        = 0x26c
	KEY_NUMERIC_12                        = 0x26d
	KEY_AUDIO_DESC                        = 0x26e
	KEY_3D_MODE                           = 0x26f
	KEY_NEXT_FAVORITE                     = 0x270
	KEY_STOP_RECORD                       = 0x271
	KEY_PAUSE_RECORD                      = 0x272
	KEY_VOD                               = 0x273
	KEY_UNMUTE                            = 0x274
	KEY_FASTREVERSE                       = 0x275
	KEY_SLOWREVERSE                       = 0x276
	KEY_DATA                              = 0x277
	KEY_ONSCREEN_KEYBOARD                 = 0x278
	BTN_TRIGGER_HAPPY                     = 0x2c0
	BTN_TRIGGER_HAPPY1                    = 0x2c0
	BTN_TRIGGER_HAPPY2                    = 0x2c1
	BTN_TRIGGER_HAPPY3                    = 0x2c2
	BTN_TRIGGER_HAPPY4                    = 0x2c3
	BTN_TRIGGER_HAPPY5                    = 0x2c4
	BTN_TRIGGER_HAPPY6                    = 0x2c5
	BTN_TRIGGER_HAPPY7                    = 0x2c6
	BTN_TRIGGER_HAPPY8                    = 0x2c7
	BTN_TRIGGER_HAPPY9                    = 0x2c8
	BTN_TRIGGER_HAPPY10                   = 0x2c9
	BTN_TRIGGER_HAPPY11                   = 0x2ca
	BTN_TRIGGER_HAPPY12                   = 0x2cb
	BTN_TRIGGER_HAPPY13                   = 0x2cc
	BTN_TRIGGER_HAPPY14                   = 0x2cd
	BTN_TRIGGER_HAPPY15                   = 0x2ce
	BTN_TRIGGER_HAPPY16                   = 0x2cf
	BTN_TRIGGER_HAPPY17                   = 0x2d0
	BTN_TRIGGER_HAPPY18                   = 0x2d1
	BTN_TRIGGER_HAPPY19                   = 0x2d2
	BTN_TRIGGER_HAPPY20                   = 0x2d3
	BTN_TRIGGER_HAPPY21                   = 0x2d4
	BTN_TRIGGER_HAPPY22                   = 0x2d5
	BTN_TRIGGER_HAPPY23                   = 0x2d6
	BTN_TRIGGER_HAPPY24                   = 0x2d7
	BTN_TRIGGER_HAPPY25                   = 0x2d8
	BTN_TRIGGER_HAPPY26                   = 0x2d9
	BTN_TRIGGER_HAPPY27                   = 0x2da
	BTN_TRIGGER_HAPPY28                   = 0x2db
	BTN_TRIGGER_HAPPY29                   = 0x2dc
	BTN_TRIGGER_HAPPY30                   = 0x2dd
	BTN_TRIGGER_HAPPY31                   = 0x2de
	BTN_TRIGGER_HAPPY32                   = 0x2df
	BTN_TRIGGER_HAPPY33                   = 0x2e0
	BTN_TRIGGER_HAPPY34                   = 0x2e1
	BTN_TRIGGER_HAPPY35                   = 0x2e2
	BTN_TRIGGER_HAPPY36                   = 0x2e3
	BTN_TRIGGER_HAPPY37                   = 0x2e4
	BTN_TRIGGER_HAPPY38                   = 0x2e5
	BTN_TRIGGER_HAPPY39                   = 0x2e6
	BTN_TRIGGER_HAPPY40                   = 0x2e7
	KEY_MIN_INTERESTING                   = KEY_MUTE
	KEY_MAX                               = 0x2ff
	KEY_CNT                               = (KEY_MAX + 1)
)

func (v KEY_CODE) String() string {
	switch v {
	case KEY_RESERVED:
		return fmt.Sprintf("KEY_RESERVED(0x%04x)", uint16(v))
	case KEY_ESC:
		return fmt.Sprintf("KEY_ESC(0x%04x)", uint16(v))
	case KEY_1:
		return fmt.Sprintf("KEY_1(0x%04x)", uint16(v))
	case KEY_2:
		return fmt.Sprintf("KEY_2(0x%04x)", uint16(v))
	case KEY_3:
		return fmt.Sprintf("KEY_3(0x%04x)", uint16(v))
	case KEY_4:
		return fmt.Sprintf("KEY_4(0x%04x)", uint16(v))
	case KEY_5:
		return fmt.Sprintf("KEY_5(0x%04x)", uint16(v))
	case KEY_6:
		return fmt.Sprintf("KEY_6(0x%04x)", uint16(v))
	case KEY_7:
		return fmt.Sprintf("KEY_7(0x%04x)", uint16(v))
	case KEY_8:
		return fmt.Sprintf("KEY_8(0x%04x)", uint16(v))
	case KEY_9:
		return fmt.Sprintf("KEY_9(0x%04x)", uint16(v))
	case KEY_0:
		return fmt.Sprintf("KEY_0(0x%04x)", uint16(v))
	case KEY_MINUS:
		return fmt.Sprintf("KEY_MINUS(0x%04x)", uint16(v))
	case KEY_EQUAL:
		return fmt.Sprintf("KEY_EQUAL(0x%04x)", uint16(v))
	case KEY_BACKSPACE:
		return fmt.Sprintf("KEY_BACKSPACE(0x%04x)", uint16(v))
	case KEY_TAB:
		return fmt.Sprintf("KEY_TAB(0x%04x)", uint16(v))
	case KEY_Q:
		return fmt.Sprintf("KEY_Q(0x%04x)", uint16(v))
	case KEY_W:
		return fmt.Sprintf("KEY_W(0x%04x)", uint16(v))
	case KEY_E:
		return fmt.Sprintf("KEY_E(0x%04x)", uint16(v))
	case KEY_R:
		return fmt.Sprintf("KEY_R(0x%04x)", uint16(v))
	case KEY_T:
		return fmt.Sprintf("KEY_T(0x%04x)", uint16(v))
	case KEY_Y:
		return fmt.Sprintf("KEY_Y(0x%04x)", uint16(v))
	case KEY_U:
		return fmt.Sprintf("KEY_U(0x%04x)", uint16(v))
	case KEY_I:
		return fmt.Sprintf("KEY_I(0x%04x)", uint16(v))
	case KEY_O:
		return fmt.Sprintf("KEY_O(0x%04x)", uint16(v))
	case KEY_P:
		return fmt.Sprintf("KEY_P(0x%04x)", uint16(v))
	case KEY_LEFTBRACE:
		return fmt.Sprintf("KEY_LEFTBRACE(0x%04x)", uint16(v))
	case KEY_RIGHTBRACE:
		return fmt.Sprintf("KEY_RIGHTBRACE(0x%04x)", uint16(v))
	case KEY_ENTER:
		return fmt.Sprintf("KEY_ENTER(0x%04x)", uint16(v))
	case KEY_LEFTCTRL:
		return fmt.Sprintf("KEY_LEFTCTRL(0x%04x)", uint16(v))
	case KEY_A:
		return fmt.Sprintf("KEY_A(0x%04x)", uint16(v))
	case KEY_S:
		return fmt.Sprintf("KEY_S(0x%04x)", uint16(v))
	case KEY_D:
		return fmt.Sprintf("KEY_D(0x%04x)", uint16(v))
	case KEY_F:
		return fmt.Sprintf("KEY_F(0x%04x)", uint16(v))
	case KEY_G:
		return fmt.Sprintf("KEY_G(0x%04x)", uint16(v))
	case KEY_H:
		return fmt.Sprintf("KEY_H(0x%04x)", uint16(v))
	case KEY_J:
		return fmt.Sprintf("KEY_J(0x%04x)", uint16(v))
	case KEY_K:
		return fmt.Sprintf("KEY_K(0x%04x)", uint16(v))
	case KEY_L:
		return fmt.Sprintf("KEY_L(0x%04x)", uint16(v))
	case KEY_SEMICOLON:
		return fmt.Sprintf("KEY_SEMICOLON(0x%04x)", uint16(v))
	case KEY_APOSTROPHE:
		return fmt.Sprintf("KEY_APOSTROPHE(0x%04x)", uint16(v))
	case KEY_GRAVE:
		return fmt.Sprintf("KEY_GRAVE(0x%04x)", uint16(v))
	case KEY_LEFTSHIFT:
		return fmt.Sprintf("KEY_LEFTSHIFT(0x%04x)", uint16(v))
	case KEY_BACKSLASH:
		return fmt.Sprintf("KEY_BACKSLASH(0x%04x)", uint16(v))
	case KEY_Z:
		return fmt.Sprintf("KEY_Z(0x%04x)", uint16(v))
	case KEY_X:
		return fmt.Sprintf("KEY_X(0x%04x)", uint16(v))
	case KEY_C:
		return fmt.Sprintf("KEY_C(0x%04x)", uint16(v))
	case KEY_V:
		return fmt.Sprintf("KEY_V(0x%04x)", uint16(v))
	case KEY_B:
		return fmt.Sprintf("KEY_B(0x%04x)", uint16(v))
	case KEY_N:
		return fmt.Sprintf("KEY_N(0x%04x)", uint16(v))
	case KEY_M:
		return fmt.Sprintf("KEY_M(0x%04x)", uint16(v))
	case KEY_COMMA:
		return fmt.Sprintf("KEY_COMMA(0x%04x)", uint16(v))
	case KEY_DOT:
		return fmt.Sprintf("KEY_DOT(0x%04x)", uint16(v))
	case KEY_SLASH:
		return fmt.Sprintf("KEY_SLASH(0x%04x)", uint16(v))
	case KEY_RIGHTSHIFT:
		return fmt.Sprintf("KEY_RIGHTSHIFT(0x%04x)", uint16(v))
	case KEY_KPASTERISK:
		return fmt.Sprintf("KEY_KPASTERISK(0x%04x)", uint16(v))
	case KEY_LEFTALT:
		return fmt.Sprintf("KEY_LEFTALT(0x%04x)", uint16(v))
	case KEY_SPACE:
		return fmt.Sprintf("KEY_SPACE(0x%04x)", uint16(v))
	case KEY_CAPSLOCK:
		return fmt.Sprintf("KEY_CAPSLOCK(0x%04x)", uint16(v))
	case KEY_F1:
		return fmt.Sprintf("KEY_F1(0x%04x)", uint16(v))
	case KEY_F2:
		return fmt.Sprintf("KEY_F2(0x%04x)", uint16(v))
	case KEY_F3:
		return fmt.Sprintf("KEY_F3(0x%04x)", uint16(v))
	case KEY_F4:
		return fmt.Sprintf("KEY_F4(0x%04x)", uint16(v))
	case KEY_F5:
		return fmt.Sprintf("KEY_F5(0x%04x)", uint16(v))
	case KEY_F6:
		return fmt.Sprintf("KEY_F6(0x%04x)", uint16(v))
	case KEY_F7:
		return fmt.Sprintf("KEY_F7(0x%04x)", uint16(v))
	case KEY_F8:
		return fmt.Sprintf("KEY_F8(0x%04x)", uint16(v))
	case KEY_F9:
		return fmt.Sprintf("KEY_F9(0x%04x)", uint16(v))
	case KEY_F10:
		return fmt.Sprintf("KEY_F10(0x%04x)", uint16(v))
	case KEY_NUMLOCK:
		return fmt.Sprintf("KEY_NUMLOCK(0x%04x)", uint16(v))
	case KEY_SCROLLLOCK:
		return fmt.Sprintf("KEY_SCROLLLOCK(0x%04x)", uint16(v))
	case KEY_KP7:
		return fmt.Sprintf("KEY_KP7(0x%04x)", uint16(v))
	case KEY_KP8:
		return fmt.Sprintf("KEY_KP8(0x%04x)", uint16(v))
	case KEY_KP9:
		return fmt.Sprintf("KEY_KP9(0x%04x)", uint16(v))
	case KEY_KPMINUS:
		return fmt.Sprintf("KEY_KPMINUS(0x%04x)", uint16(v))
	case KEY_KP4:
		return fmt.Sprintf("KEY_KP4(0x%04x)", uint16(v))
	case KEY_KP5:
		return fmt.Sprintf("KEY_KP5(0x%04x)", uint16(v))
	case KEY_KP6:
		return fmt.Sprintf("KEY_KP6(0x%04x)", uint16(v))
	case KEY_KPPLUS:
		return fmt.Sprintf("KEY_KPPLUS(0x%04x)", uint16(v))
	case KEY_KP1:
		return fmt.Sprintf("KEY_KP1(0x%04x)", uint16(v))
	case KEY_KP2:
		return fmt.Sprintf("KEY_KP2(0x%04x)", uint16(v))
	case KEY_KP3:
		return fmt.Sprintf("KEY_KP3(0x%04x)", uint16(v))
	case KEY_KP0:
		return fmt.Sprintf("KEY_KP0(0x%04x)", uint16(v))
	case KEY_KPDOT:
		return fmt.Sprintf("KEY_KPDOT(0x%04x)", uint16(v))
	case KEY_ZENKAKUHANKAKU:
		return fmt.Sprintf("KEY_ZENKAKUHANKAKU(0x%04x)", uint16(v))
	case KEY_102ND:
		return fmt.Sprintf("KEY_102ND(0x%04x)", uint16(v))
	case KEY_F11:
		return fmt.Sprintf("KEY_F11(0x%04x)", uint16(v))
	case KEY_F12:
		return fmt.Sprintf("KEY_F12(0x%04x)", uint16(v))
	case KEY_RO:
		return fmt.Sprintf("KEY_RO(0x%04x)", uint16(v))
	case KEY_KATAKANA:
		return fmt.Sprintf("KEY_KATAKANA(0x%04x)", uint16(v))
	case KEY_HIRAGANA:
		return fmt.Sprintf("KEY_HIRAGANA(0x%04x)", uint16(v))
	case KEY_HENKAN:
		return fmt.Sprintf("KEY_HENKAN(0x%04x)", uint16(v))
	case KEY_KATAKANAHIRAGANA:
		return fmt.Sprintf("KEY_KATAKANAHIRAGANA(0x%04x)", uint16(v))
	case KEY_MUHENKAN:
		return fmt.Sprintf("KEY_MUHENKAN(0x%04x)", uint16(v))
	case KEY_KPJPCOMMA:
		return fmt.Sprintf("KEY_KPJPCOMMA(0x%04x)", uint16(v))
	case KEY_KPENTER:
		return fmt.Sprintf("KEY_KPENTER(0x%04x)", uint16(v))
	case KEY_RIGHTCTRL:
		return fmt.Sprintf("KEY_RIGHTCTRL(0x%04x)", uint16(v))
	case KEY_KPSLASH:
		return fmt.Sprintf("KEY_KPSLASH(0x%04x)", uint16(v))
	case KEY_SYSRQ:
		return fmt.Sprintf("KEY_SYSRQ(0x%04x)", uint16(v))
	case KEY_RIGHTALT:
		return fmt.Sprintf("KEY_RIGHTALT(0x%04x)", uint16(v))
	case KEY_LINEFEED:
		return fmt.Sprintf("KEY_LINEFEED(0x%04x)", uint16(v))
	case KEY_HOME:
		return fmt.Sprintf("KEY_HOME(0x%04x)", uint16(v))
	case KEY_UP:
		return fmt.Sprintf("KEY_UP(0x%04x)", uint16(v))
	case KEY_PAGEUP:
		return fmt.Sprintf("KEY_PAGEUP(0x%04x)", uint16(v))
	case KEY_LEFT:
		return fmt.Sprintf("KEY_LEFT(0x%04x)", uint16(v))
	case KEY_RIGHT:
		return fmt.Sprintf("KEY_RIGHT(0x%04x)", uint16(v))
	case KEY_END:
		return fmt.Sprintf("KEY_END(0x%04x)", uint16(v))
	case KEY_DOWN:
		return fmt.Sprintf("KEY_DOWN(0x%04x)", uint16(v))
	case KEY_PAGEDOWN:
		return fmt.Sprintf("KEY_PAGEDOWN(0x%04x)", uint16(v))
	case KEY_INSERT:
		return fmt.Sprintf("KEY_INSERT(0x%04x)", uint16(v))
	case KEY_DELETE:
		return fmt.Sprintf("KEY_DELETE(0x%04x)", uint16(v))
	case KEY_MACRO:
		return fmt.Sprintf("KEY_MACRO(0x%04x)", uint16(v))
	case KEY_VOLUMEDOWN:
		return fmt.Sprintf("KEY_VOLUMEDOWN(0x%04x)", uint16(v))
	case KEY_VOLUMEUP:
		return fmt.Sprintf("KEY_VOLUMEUP(0x%04x)", uint16(v))
	case KEY_POWER:
		return fmt.Sprintf("KEY_POWER(0x%04x)", uint16(v))
	case KEY_KPEQUAL:
		return fmt.Sprintf("KEY_KPEQUAL(0x%04x)", uint16(v))
	case KEY_KPPLUSMINUS:
		return fmt.Sprintf("KEY_KPPLUSMINUS(0x%04x)", uint16(v))
	case KEY_PAUSE:
		return fmt.Sprintf("KEY_PAUSE(0x%04x)", uint16(v))
	case KEY_SCALE:
		return fmt.Sprintf("KEY_SCALE(0x%04x)", uint16(v))
	case KEY_KPCOMMA:
		return fmt.Sprintf("KEY_KPCOMMA(0x%04x)", uint16(v))
	case KEY_HANGEUL: // | KEY_HANGUEL:
		return fmt.Sprintf("KEY_HANGEUL|KEY_HANGUEL(0x%04x)", uint16(v))
	case KEY_HANJA:
		return fmt.Sprintf("KEY_HANJA(0x%04x)", uint16(v))
	case KEY_YEN:
		return fmt.Sprintf("KEY_YEN(0x%04x)", uint16(v))
	case KEY_LEFTMETA:
		return fmt.Sprintf("KEY_LEFTMETA(0x%04x)", uint16(v))
	case KEY_RIGHTMETA:
		return fmt.Sprintf("KEY_RIGHTMETA(0x%04x)", uint16(v))
	case KEY_COMPOSE:
		return fmt.Sprintf("KEY_COMPOSE(0x%04x)", uint16(v))
	case KEY_STOP:
		return fmt.Sprintf("KEY_STOP(0x%04x)", uint16(v))
	case KEY_AGAIN:
		return fmt.Sprintf("KEY_AGAIN(0x%04x)", uint16(v))
	case KEY_PROPS:
		return fmt.Sprintf("KEY_PROPS(0x%04x)", uint16(v))
	case KEY_UNDO:
		return fmt.Sprintf("KEY_UNDO(0x%04x)", uint16(v))
	case KEY_FRONT:
		return fmt.Sprintf("KEY_FRONT(0x%04x)", uint16(v))
	case KEY_COPY:
		return fmt.Sprintf("KEY_COPY(0x%04x)", uint16(v))
	case KEY_OPEN:
		return fmt.Sprintf("KEY_OPEN(0x%04x)", uint16(v))
	case KEY_PASTE:
		return fmt.Sprintf("KEY_PASTE(0x%04x)", uint16(v))
	case KEY_FIND:
		return fmt.Sprintf("KEY_FIND(0x%04x)", uint16(v))
	case KEY_CUT:
		return fmt.Sprintf("KEY_CUT(0x%04x)", uint16(v))
	case KEY_HELP:
		return fmt.Sprintf("KEY_HELP(0x%04x)", uint16(v))
	case KEY_MENU:
		return fmt.Sprintf("KEY_MENU(0x%04x)", uint16(v))
	case KEY_CALC:
		return fmt.Sprintf("KEY_CALC(0x%04x)", uint16(v))
	case KEY_SETUP:
		return fmt.Sprintf("KEY_SETUP(0x%04x)", uint16(v))
	case KEY_SLEEP:
		return fmt.Sprintf("KEY_SLEEP(0x%04x)", uint16(v))
	case KEY_WAKEUP:
		return fmt.Sprintf("KEY_WAKEUP(0x%04x)", uint16(v))
	case KEY_FILE:
		return fmt.Sprintf("KEY_FILE(0x%04x)", uint16(v))
	case KEY_SENDFILE:
		return fmt.Sprintf("KEY_SENDFILE(0x%04x)", uint16(v))
	case KEY_DELETEFILE:
		return fmt.Sprintf("KEY_DELETEFILE(0x%04x)", uint16(v))
	case KEY_XFER:
		return fmt.Sprintf("KEY_XFER(0x%04x)", uint16(v))
	case KEY_PROG1:
		return fmt.Sprintf("KEY_PROG1(0x%04x)", uint16(v))
	case KEY_PROG2:
		return fmt.Sprintf("KEY_PROG2(0x%04x)", uint16(v))
	case KEY_WWW:
		return fmt.Sprintf("KEY_WWW(0x%04x)", uint16(v))
	case KEY_MSDOS:
		return fmt.Sprintf("KEY_MSDOS(0x%04x)", uint16(v))
	case KEY_COFFEE: // | KEY_SCREENLOCK:
		return fmt.Sprintf("KEY_COFFEE|KEY_SCREENLOCK(0x%04x)", uint16(v))
	case KEY_ROTATE_DISPLAY: // | KEY_DIRECTION:
		return fmt.Sprintf("KEY_ROTATE_DISPLAY|KEY_DIRECTION(0x%04x)", uint16(v))
	case KEY_CYCLEWINDOWS:
		return fmt.Sprintf("KEY_CYCLEWINDOWS(0x%04x)", uint16(v))
	case KEY_MAIL:
		return fmt.Sprintf("KEY_MAIL(0x%04x)", uint16(v))
	case KEY_BOOKMARKS:
		return fmt.Sprintf("KEY_BOOKMARKS(0x%04x)", uint16(v))
	case KEY_COMPUTER:
		return fmt.Sprintf("KEY_COMPUTER(0x%04x)", uint16(v))
	case KEY_BACK:
		return fmt.Sprintf("KEY_BACK(0x%04x)", uint16(v))
	case KEY_FORWARD:
		return fmt.Sprintf("KEY_FORWARD(0x%04x)", uint16(v))
	case KEY_CLOSECD:
		return fmt.Sprintf("KEY_CLOSECD(0x%04x)", uint16(v))
	case KEY_EJECTCD:
		return fmt.Sprintf("KEY_EJECTCD(0x%04x)", uint16(v))
	case KEY_EJECTCLOSECD:
		return fmt.Sprintf("KEY_EJECTCLOSECD(0x%04x)", uint16(v))
	case KEY_NEXTSONG:
		return fmt.Sprintf("KEY_NEXTSONG(0x%04x)", uint16(v))
	case KEY_PLAYPAUSE:
		return fmt.Sprintf("KEY_PLAYPAUSE(0x%04x)", uint16(v))
	case KEY_PREVIOUSSONG:
		return fmt.Sprintf("KEY_PREVIOUSSONG(0x%04x)", uint16(v))
	case KEY_STOPCD:
		return fmt.Sprintf("KEY_STOPCD(0x%04x)", uint16(v))
	case KEY_RECORD:
		return fmt.Sprintf("KEY_RECORD(0x%04x)", uint16(v))
	case KEY_REWIND:
		return fmt.Sprintf("KEY_REWIND(0x%04x)", uint16(v))
	case KEY_PHONE:
		return fmt.Sprintf("KEY_PHONE(0x%04x)", uint16(v))
	case KEY_ISO:
		return fmt.Sprintf("KEY_ISO(0x%04x)", uint16(v))
	case KEY_CONFIG:
		return fmt.Sprintf("KEY_CONFIG(0x%04x)", uint16(v))
	case KEY_HOMEPAGE:
		return fmt.Sprintf("KEY_HOMEPAGE(0x%04x)", uint16(v))
	case KEY_REFRESH:
		return fmt.Sprintf("KEY_REFRESH(0x%04x)", uint16(v))
	case KEY_EXIT:
		return fmt.Sprintf("KEY_EXIT(0x%04x)", uint16(v))
	case KEY_MOVE:
		return fmt.Sprintf("KEY_MOVE(0x%04x)", uint16(v))
	case KEY_EDIT:
		return fmt.Sprintf("KEY_EDIT(0x%04x)", uint16(v))
	case KEY_SCROLLUP:
		return fmt.Sprintf("KEY_SCROLLUP(0x%04x)", uint16(v))
	case KEY_SCROLLDOWN:
		return fmt.Sprintf("KEY_SCROLLDOWN(0x%04x)", uint16(v))
	case KEY_KPLEFTPAREN:
		return fmt.Sprintf("KEY_KPLEFTPAREN(0x%04x)", uint16(v))
	case KEY_KPRIGHTPAREN:
		return fmt.Sprintf("KEY_KPRIGHTPAREN(0x%04x)", uint16(v))
	case KEY_NEW:
		return fmt.Sprintf("KEY_NEW(0x%04x)", uint16(v))
	case KEY_REDO:
		return fmt.Sprintf("KEY_REDO(0x%04x)", uint16(v))
	case KEY_F13:
		return fmt.Sprintf("KEY_F13(0x%04x)", uint16(v))
	case KEY_F14:
		return fmt.Sprintf("KEY_F14(0x%04x)", uint16(v))
	case KEY_F15:
		return fmt.Sprintf("KEY_F15(0x%04x)", uint16(v))
	case KEY_F16:
		return fmt.Sprintf("KEY_F16(0x%04x)", uint16(v))
	case KEY_F17:
		return fmt.Sprintf("KEY_F17(0x%04x)", uint16(v))
	case KEY_F18:
		return fmt.Sprintf("KEY_F18(0x%04x)", uint16(v))
	case KEY_F19:
		return fmt.Sprintf("KEY_F19(0x%04x)", uint16(v))
	case KEY_F20:
		return fmt.Sprintf("KEY_F20(0x%04x)", uint16(v))
	case KEY_F21:
		return fmt.Sprintf("KEY_F21(0x%04x)", uint16(v))
	case KEY_F22:
		return fmt.Sprintf("KEY_F22(0x%04x)", uint16(v))
	case KEY_F23:
		return fmt.Sprintf("KEY_F23(0x%04x)", uint16(v))
	case KEY_F24:
		return fmt.Sprintf("KEY_F24(0x%04x)", uint16(v))
	case KEY_PLAYCD:
		return fmt.Sprintf("KEY_PLAYCD(0x%04x)", uint16(v))
	case KEY_PAUSECD:
		return fmt.Sprintf("KEY_PAUSECD(0x%04x)", uint16(v))
	case KEY_PROG3:
		return fmt.Sprintf("KEY_PROG3(0x%04x)", uint16(v))
	case KEY_PROG4:
		return fmt.Sprintf("KEY_PROG4(0x%04x)", uint16(v))
	case KEY_DASHBOARD:
		return fmt.Sprintf("KEY_DASHBOARD(0x%04x)", uint16(v))
	case KEY_SUSPEND:
		return fmt.Sprintf("KEY_SUSPEND(0x%04x)", uint16(v))
	case KEY_CLOSE:
		return fmt.Sprintf("KEY_CLOSE(0x%04x)", uint16(v))
	case KEY_PLAY:
		return fmt.Sprintf("KEY_PLAY(0x%04x)", uint16(v))
	case KEY_FASTFORWARD:
		return fmt.Sprintf("KEY_FASTFORWARD(0x%04x)", uint16(v))
	case KEY_BASSBOOST:
		return fmt.Sprintf("KEY_BASSBOOST(0x%04x)", uint16(v))
	case KEY_PRINT:
		return fmt.Sprintf("KEY_PRINT(0x%04x)", uint16(v))
	case KEY_HP:
		return fmt.Sprintf("KEY_HP(0x%04x)", uint16(v))
	case KEY_CAMERA:
		return fmt.Sprintf("KEY_CAMERA(0x%04x)", uint16(v))
	case KEY_SOUND:
		return fmt.Sprintf("KEY_SOUND(0x%04x)", uint16(v))
	case KEY_QUESTION:
		return fmt.Sprintf("KEY_QUESTION(0x%04x)", uint16(v))
	case KEY_EMAIL:
		return fmt.Sprintf("KEY_EMAIL(0x%04x)", uint16(v))
	case KEY_CHAT:
		return fmt.Sprintf("KEY_CHAT(0x%04x)", uint16(v))
	case KEY_SEARCH:
		return fmt.Sprintf("KEY_SEARCH(0x%04x)", uint16(v))
	case KEY_CONNECT:
		return fmt.Sprintf("KEY_CONNECT(0x%04x)", uint16(v))
	case KEY_FINANCE:
		return fmt.Sprintf("KEY_FINANCE(0x%04x)", uint16(v))
	case KEY_SPORT:
		return fmt.Sprintf("KEY_SPORT(0x%04x)", uint16(v))
	case KEY_SHOP:
		return fmt.Sprintf("KEY_SHOP(0x%04x)", uint16(v))
	case KEY_ALTERASE:
		return fmt.Sprintf("KEY_ALTERASE(0x%04x)", uint16(v))
	case KEY_CANCEL:
		return fmt.Sprintf("KEY_CANCEL(0x%04x)", uint16(v))
	case KEY_BRIGHTNESSDOWN:
		return fmt.Sprintf("KEY_BRIGHTNESSDOWN(0x%04x)", uint16(v))
	case KEY_BRIGHTNESSUP:
		return fmt.Sprintf("KEY_BRIGHTNESSUP(0x%04x)", uint16(v))
	case KEY_MEDIA:
		return fmt.Sprintf("KEY_MEDIA(0x%04x)", uint16(v))
	case KEY_SWITCHVIDEOMODE:
		return fmt.Sprintf("KEY_SWITCHVIDEOMODE(0x%04x)", uint16(v))
	case KEY_KBDILLUMTOGGLE:
		return fmt.Sprintf("KEY_KBDILLUMTOGGLE(0x%04x)", uint16(v))
	case KEY_KBDILLUMDOWN:
		return fmt.Sprintf("KEY_KBDILLUMDOWN(0x%04x)", uint16(v))
	case KEY_KBDILLUMUP:
		return fmt.Sprintf("KEY_KBDILLUMUP(0x%04x)", uint16(v))
	case KEY_SEND:
		return fmt.Sprintf("KEY_SEND(0x%04x)", uint16(v))
	case KEY_REPLY:
		return fmt.Sprintf("KEY_REPLY(0x%04x)", uint16(v))
	case KEY_FORWARDMAIL:
		return fmt.Sprintf("KEY_FORWARDMAIL(0x%04x)", uint16(v))
	case KEY_SAVE:
		return fmt.Sprintf("KEY_SAVE(0x%04x)", uint16(v))
	case KEY_DOCUMENTS:
		return fmt.Sprintf("KEY_DOCUMENTS(0x%04x)", uint16(v))
	case KEY_BATTERY:
		return fmt.Sprintf("KEY_BATTERY(0x%04x)", uint16(v))
	case KEY_BLUETOOTH:
		return fmt.Sprintf("KEY_BLUETOOTH(0x%04x)", uint16(v))
	case KEY_WLAN:
		return fmt.Sprintf("KEY_WLAN(0x%04x)", uint16(v))
	case KEY_UWB:
		return fmt.Sprintf("KEY_UWB(0x%04x)", uint16(v))
	case KEY_UNKNOWN:
		return fmt.Sprintf("KEY_UNKNOWN(0x%04x)", uint16(v))
	case KEY_VIDEO_NEXT:
		return fmt.Sprintf("KEY_VIDEO_NEXT(0x%04x)", uint16(v))
	case KEY_VIDEO_PREV:
		return fmt.Sprintf("KEY_VIDEO_PREV(0x%04x)", uint16(v))
	case KEY_BRIGHTNESS_CYCLE:
		return fmt.Sprintf("KEY_BRIGHTNESS_CYCLE(0x%04x)", uint16(v))
	case KEY_BRIGHTNESS_AUTO: // | KEY_BRIGHTNESS_ZERO:
		return fmt.Sprintf("KEY_BRIGHTNESS_AUTO|KEY_BRIGHTNESS_ZERO(0x%04x)", uint16(v))
	case KEY_DISPLAY_OFF:
		return fmt.Sprintf("KEY_DISPLAY_OFF(0x%04x)", uint16(v))
	case KEY_WWAN: // | KEY_WIMAX:
		return fmt.Sprintf("KEY_WWAN|KEY_WIMAX(0x%04x)", uint16(v))
	case KEY_RFKILL:
		return fmt.Sprintf("KEY_RFKILL(0x%04x)", uint16(v))
	case KEY_MICMUTE:
		return fmt.Sprintf("KEY_MICMUTE(0x%04x)", uint16(v))
	case BTN_MISC: // | BTN_0:
		return fmt.Sprintf("BTN_MISC|BTN_0(0x%04x)", uint16(v))
	case BTN_1:
		return fmt.Sprintf("BTN_1(0x%04x)", uint16(v))
	case BTN_2:
		return fmt.Sprintf("BTN_2(0x%04x)", uint16(v))
	case BTN_3:
		return fmt.Sprintf("BTN_3(0x%04x)", uint16(v))
	case BTN_4:
		return fmt.Sprintf("BTN_4(0x%04x)", uint16(v))
	case BTN_5:
		return fmt.Sprintf("BTN_5(0x%04x)", uint16(v))
	case BTN_6:
		return fmt.Sprintf("BTN_6(0x%04x)", uint16(v))
	case BTN_7:
		return fmt.Sprintf("BTN_7(0x%04x)", uint16(v))
	case BTN_8:
		return fmt.Sprintf("BTN_8(0x%04x)", uint16(v))
	case BTN_9:
		return fmt.Sprintf("BTN_9(0x%04x)", uint16(v))
	case BTN_MOUSE: // | BTN_LEFT:
		return fmt.Sprintf("BTN_MOUSE|BTN_LEFT(0x%04x)", uint16(v))
	case BTN_RIGHT:
		return fmt.Sprintf("BTN_RIGHT(0x%04x)", uint16(v))
	case BTN_MIDDLE:
		return fmt.Sprintf("BTN_MIDDLE(0x%04x)", uint16(v))
	case BTN_SIDE:
		return fmt.Sprintf("BTN_SIDE(0x%04x)", uint16(v))
	case BTN_EXTRA:
		return fmt.Sprintf("BTN_EXTRA(0x%04x)", uint16(v))
	case BTN_FORWARD:
		return fmt.Sprintf("BTN_FORWARD(0x%04x)", uint16(v))
	case BTN_BACK:
		return fmt.Sprintf("BTN_BACK(0x%04x)", uint16(v))
	case BTN_TASK:
		return fmt.Sprintf("BTN_TASK(0x%04x)", uint16(v))
	case BTN_JOYSTICK: // | BTN_TRIGGER:
		return fmt.Sprintf("BTN_JOYSTICK|BTN_TRIGGER(0x%04x)", uint16(v))
	case BTN_THUMB:
		return fmt.Sprintf("BTN_THUMB(0x%04x)", uint16(v))
	case BTN_THUMB2:
		return fmt.Sprintf("BTN_THUMB2(0x%04x)", uint16(v))
	case BTN_TOP:
		return fmt.Sprintf("BTN_TOP(0x%04x)", uint16(v))
	case BTN_TOP2:
		return fmt.Sprintf("BTN_TOP2(0x%04x)", uint16(v))
	case BTN_PINKIE:
		return fmt.Sprintf("BTN_PINKIE(0x%04x)", uint16(v))
	case BTN_BASE:
		return fmt.Sprintf("BTN_BASE(0x%04x)", uint16(v))
	case BTN_BASE2:
		return fmt.Sprintf("BTN_BASE2(0x%04x)", uint16(v))
	case BTN_BASE3:
		return fmt.Sprintf("BTN_BASE3(0x%04x)", uint16(v))
	case BTN_BASE4:
		return fmt.Sprintf("BTN_BASE4(0x%04x)", uint16(v))
	case BTN_BASE5:
		return fmt.Sprintf("BTN_BASE5(0x%04x)", uint16(v))
	case BTN_BASE6:
		return fmt.Sprintf("BTN_BASE6(0x%04x)", uint16(v))
	case BTN_DEAD:
		return fmt.Sprintf("BTN_DEAD(0x%04x)", uint16(v))
	case BTN_GAMEPAD: // | BTN_SOUTH: // | BTN_A:
		return fmt.Sprintf("BTN_GAMEPAD|BTN_SOUTH|BTN_A(0x%04x)", uint16(v))
	case BTN_EAST: // | BTN_B:
		return fmt.Sprintf("BTN_EAST|BTN_B(0x%04x)", uint16(v))
	case BTN_C:
		return fmt.Sprintf("BTN_C(0x%04x)", uint16(v))
	case BTN_NORTH: // | BTN_X:
		return fmt.Sprintf("BTN_NORTH|BTN_X(0x%04x)", uint16(v))
	case BTN_WEST: // | BTN_Y:
		return fmt.Sprintf("BTN_WEST|BTN_Y(0x%04x)", uint16(v))
	case BTN_Z:
		return fmt.Sprintf("BTN_Z(0x%04x)", uint16(v))
	case BTN_TL:
		return fmt.Sprintf("BTN_TL(0x%04x)", uint16(v))
	case BTN_TR:
		return fmt.Sprintf("BTN_TR(0x%04x)", uint16(v))
	case BTN_TL2:
		return fmt.Sprintf("BTN_TL2(0x%04x)", uint16(v))
	case BTN_TR2:
		return fmt.Sprintf("BTN_TR2(0x%04x)", uint16(v))
	case BTN_SELECT:
		return fmt.Sprintf("BTN_SELECT(0x%04x)", uint16(v))
	case BTN_START:
		return fmt.Sprintf("BTN_START(0x%04x)", uint16(v))
	case BTN_MODE:
		return fmt.Sprintf("BTN_MODE(0x%04x)", uint16(v))
	case BTN_THUMBL:
		return fmt.Sprintf("BTN_THUMBL(0x%04x)", uint16(v))
	case BTN_THUMBR:
		return fmt.Sprintf("BTN_THUMBR(0x%04x)", uint16(v))
	case BTN_DIGI: // | BTN_TOOL_PEN:
		return fmt.Sprintf("BTN_DIGI|BTN_TOOL_PEN(0x%04x)", uint16(v))
	case BTN_TOOL_RUBBER:
		return fmt.Sprintf("BTN_TOOL_RUBBER(0x%04x)", uint16(v))
	case BTN_TOOL_BRUSH:
		return fmt.Sprintf("BTN_TOOL_BRUSH(0x%04x)", uint16(v))
	case BTN_TOOL_PENCIL:
		return fmt.Sprintf("BTN_TOOL_PENCIL(0x%04x)", uint16(v))
	case BTN_TOOL_AIRBRUSH:
		return fmt.Sprintf("BTN_TOOL_AIRBRUSH(0x%04x)", uint16(v))
	case BTN_TOOL_FINGER:
		return fmt.Sprintf("BTN_TOOL_FINGER(0x%04x)", uint16(v))
	case BTN_TOOL_MOUSE:
		return fmt.Sprintf("BTN_TOOL_MOUSE(0x%04x)", uint16(v))
	case BTN_TOOL_LENS:
		return fmt.Sprintf("BTN_TOOL_LENS(0x%04x)", uint16(v))
	case BTN_TOOL_QUINTTAP:
		return fmt.Sprintf("BTN_TOOL_QUINTTAP(0x%04x)", uint16(v))
	case BTN_TOUCH:
		return fmt.Sprintf("BTN_TOUCH(0x%04x)", uint16(v))
	case BTN_STYLUS:
		return fmt.Sprintf("BTN_STYLUS(0x%04x)", uint16(v))
	case BTN_STYLUS2:
		return fmt.Sprintf("BTN_STYLUS2(0x%04x)", uint16(v))
	case BTN_TOOL_DOUBLETAP:
		return fmt.Sprintf("BTN_TOOL_DOUBLETAP(0x%04x)", uint16(v))
	case BTN_TOOL_TRIPLETAP:
		return fmt.Sprintf("BTN_TOOL_TRIPLETAP(0x%04x)", uint16(v))
	case BTN_TOOL_QUADTAP:
		return fmt.Sprintf("BTN_TOOL_QUADTAP(0x%04x)", uint16(v))
	case BTN_WHEEL: // | BTN_GEAR_DOWN:
		return fmt.Sprintf("BTN_WHEEL|BTN_GEAR_DOWN(0x%04x)", uint16(v))
	case BTN_GEAR_UP:
		return fmt.Sprintf("BTN_GEAR_UP(0x%04x)", uint16(v))
	case KEY_OK:
		return fmt.Sprintf("KEY_OK(0x%04x)", uint16(v))
	case KEY_SELECT:
		return fmt.Sprintf("KEY_SELECT(0x%04x)", uint16(v))
	case KEY_GOTO:
		return fmt.Sprintf("KEY_GOTO(0x%04x)", uint16(v))
	case KEY_CLEAR:
		return fmt.Sprintf("KEY_CLEAR(0x%04x)", uint16(v))
	case KEY_POWER2:
		return fmt.Sprintf("KEY_POWER2(0x%04x)", uint16(v))
	case KEY_OPTION:
		return fmt.Sprintf("KEY_OPTION(0x%04x)", uint16(v))
	case KEY_INFO:
		return fmt.Sprintf("KEY_INFO(0x%04x)", uint16(v))
	case KEY_TIME:
		return fmt.Sprintf("KEY_TIME(0x%04x)", uint16(v))
	case KEY_VENDOR:
		return fmt.Sprintf("KEY_VENDOR(0x%04x)", uint16(v))
	case KEY_ARCHIVE:
		return fmt.Sprintf("KEY_ARCHIVE(0x%04x)", uint16(v))
	case KEY_PROGRAM:
		return fmt.Sprintf("KEY_PROGRAM(0x%04x)", uint16(v))
	case KEY_CHANNEL:
		return fmt.Sprintf("KEY_CHANNEL(0x%04x)", uint16(v))
	case KEY_FAVORITES:
		return fmt.Sprintf("KEY_FAVORITES(0x%04x)", uint16(v))
	case KEY_EPG:
		return fmt.Sprintf("KEY_EPG(0x%04x)", uint16(v))
	case KEY_PVR:
		return fmt.Sprintf("KEY_PVR(0x%04x)", uint16(v))
	case KEY_MHP:
		return fmt.Sprintf("KEY_MHP(0x%04x)", uint16(v))
	case KEY_LANGUAGE:
		return fmt.Sprintf("KEY_LANGUAGE(0x%04x)", uint16(v))
	case KEY_TITLE:
		return fmt.Sprintf("KEY_TITLE(0x%04x)", uint16(v))
	case KEY_SUBTITLE:
		return fmt.Sprintf("KEY_SUBTITLE(0x%04x)", uint16(v))
	case KEY_ANGLE:
		return fmt.Sprintf("KEY_ANGLE(0x%04x)", uint16(v))
	case KEY_ZOOM:
		return fmt.Sprintf("KEY_ZOOM(0x%04x)", uint16(v))
	case KEY_MODE:
		return fmt.Sprintf("KEY_MODE(0x%04x)", uint16(v))
	case KEY_KEYBOARD:
		return fmt.Sprintf("KEY_KEYBOARD(0x%04x)", uint16(v))
	case KEY_SCREEN:
		return fmt.Sprintf("KEY_SCREEN(0x%04x)", uint16(v))
	case KEY_PC:
		return fmt.Sprintf("KEY_PC(0x%04x)", uint16(v))
	case KEY_TV:
		return fmt.Sprintf("KEY_TV(0x%04x)", uint16(v))
	case KEY_TV2:
		return fmt.Sprintf("KEY_TV2(0x%04x)", uint16(v))
	case KEY_VCR:
		return fmt.Sprintf("KEY_VCR(0x%04x)", uint16(v))
	case KEY_VCR2:
		return fmt.Sprintf("KEY_VCR2(0x%04x)", uint16(v))
	case KEY_SAT:
		return fmt.Sprintf("KEY_SAT(0x%04x)", uint16(v))
	case KEY_SAT2:
		return fmt.Sprintf("KEY_SAT2(0x%04x)", uint16(v))
	case KEY_CD:
		return fmt.Sprintf("KEY_CD(0x%04x)", uint16(v))
	case KEY_TAPE:
		return fmt.Sprintf("KEY_TAPE(0x%04x)", uint16(v))
	case KEY_RADIO:
		return fmt.Sprintf("KEY_RADIO(0x%04x)", uint16(v))
	case KEY_TUNER:
		return fmt.Sprintf("KEY_TUNER(0x%04x)", uint16(v))
	case KEY_PLAYER:
		return fmt.Sprintf("KEY_PLAYER(0x%04x)", uint16(v))
	case KEY_TEXT:
		return fmt.Sprintf("KEY_TEXT(0x%04x)", uint16(v))
	case KEY_DVD:
		return fmt.Sprintf("KEY_DVD(0x%04x)", uint16(v))
	case KEY_AUX:
		return fmt.Sprintf("KEY_AUX(0x%04x)", uint16(v))
	case KEY_MP3:
		return fmt.Sprintf("KEY_MP3(0x%04x)", uint16(v))
	case KEY_AUDIO:
		return fmt.Sprintf("KEY_AUDIO(0x%04x)", uint16(v))
	case KEY_VIDEO:
		return fmt.Sprintf("KEY_VIDEO(0x%04x)", uint16(v))
	case KEY_DIRECTORY:
		return fmt.Sprintf("KEY_DIRECTORY(0x%04x)", uint16(v))
	case KEY_LIST:
		return fmt.Sprintf("KEY_LIST(0x%04x)", uint16(v))
	case KEY_MEMO:
		return fmt.Sprintf("KEY_MEMO(0x%04x)", uint16(v))
	case KEY_CALENDAR:
		return fmt.Sprintf("KEY_CALENDAR(0x%04x)", uint16(v))
	case KEY_RED:
		return fmt.Sprintf("KEY_RED(0x%04x)", uint16(v))
	case KEY_GREEN:
		return fmt.Sprintf("KEY_GREEN(0x%04x)", uint16(v))
	case KEY_YELLOW:
		return fmt.Sprintf("KEY_YELLOW(0x%04x)", uint16(v))
	case KEY_BLUE:
		return fmt.Sprintf("KEY_BLUE(0x%04x)", uint16(v))
	case KEY_CHANNELUP:
		return fmt.Sprintf("KEY_CHANNELUP(0x%04x)", uint16(v))
	case KEY_CHANNELDOWN:
		return fmt.Sprintf("KEY_CHANNELDOWN(0x%04x)", uint16(v))
	case KEY_FIRST:
		return fmt.Sprintf("KEY_FIRST(0x%04x)", uint16(v))
	case KEY_LAST:
		return fmt.Sprintf("KEY_LAST(0x%04x)", uint16(v))
	case KEY_AB:
		return fmt.Sprintf("KEY_AB(0x%04x)", uint16(v))
	case KEY_NEXT:
		return fmt.Sprintf("KEY_NEXT(0x%04x)", uint16(v))
	case KEY_RESTART:
		return fmt.Sprintf("KEY_RESTART(0x%04x)", uint16(v))
	case KEY_SLOW:
		return fmt.Sprintf("KEY_SLOW(0x%04x)", uint16(v))
	case KEY_SHUFFLE:
		return fmt.Sprintf("KEY_SHUFFLE(0x%04x)", uint16(v))
	case KEY_BREAK:
		return fmt.Sprintf("KEY_BREAK(0x%04x)", uint16(v))
	case KEY_PREVIOUS:
		return fmt.Sprintf("KEY_PREVIOUS(0x%04x)", uint16(v))
	case KEY_DIGITS:
		return fmt.Sprintf("KEY_DIGITS(0x%04x)", uint16(v))
	case KEY_TEEN:
		return fmt.Sprintf("KEY_TEEN(0x%04x)", uint16(v))
	case KEY_TWEN:
		return fmt.Sprintf("KEY_TWEN(0x%04x)", uint16(v))
	case KEY_VIDEOPHONE:
		return fmt.Sprintf("KEY_VIDEOPHONE(0x%04x)", uint16(v))
	case KEY_GAMES:
		return fmt.Sprintf("KEY_GAMES(0x%04x)", uint16(v))
	case KEY_ZOOMIN:
		return fmt.Sprintf("KEY_ZOOMIN(0x%04x)", uint16(v))
	case KEY_ZOOMOUT:
		return fmt.Sprintf("KEY_ZOOMOUT(0x%04x)", uint16(v))
	case KEY_ZOOMRESET:
		return fmt.Sprintf("KEY_ZOOMRESET(0x%04x)", uint16(v))
	case KEY_WORDPROCESSOR:
		return fmt.Sprintf("KEY_WORDPROCESSOR(0x%04x)", uint16(v))
	case KEY_EDITOR:
		return fmt.Sprintf("KEY_EDITOR(0x%04x)", uint16(v))
	case KEY_SPREADSHEET:
		return fmt.Sprintf("KEY_SPREADSHEET(0x%04x)", uint16(v))
	case KEY_GRAPHICSEDITOR:
		return fmt.Sprintf("KEY_GRAPHICSEDITOR(0x%04x)", uint16(v))
	case KEY_PRESENTATION:
		return fmt.Sprintf("KEY_PRESENTATION(0x%04x)", uint16(v))
	case KEY_DATABASE:
		return fmt.Sprintf("KEY_DATABASE(0x%04x)", uint16(v))
	case KEY_NEWS:
		return fmt.Sprintf("KEY_NEWS(0x%04x)", uint16(v))
	case KEY_VOICEMAIL:
		return fmt.Sprintf("KEY_VOICEMAIL(0x%04x)", uint16(v))
	case KEY_ADDRESSBOOK:
		return fmt.Sprintf("KEY_ADDRESSBOOK(0x%04x)", uint16(v))
	case KEY_MESSENGER:
		return fmt.Sprintf("KEY_MESSENGER(0x%04x)", uint16(v))
	case KEY_DISPLAYTOGGLE: // | KEY_BRIGHTNESS_TOGGLE:
		return fmt.Sprintf("KEY_DISPLAYTOGGLE|KEY_BRIGHTNESS_TOGGLE(0x%04x)", uint16(v))
	case KEY_SPELLCHECK:
		return fmt.Sprintf("KEY_SPELLCHECK(0x%04x)", uint16(v))
	case KEY_LOGOFF:
		return fmt.Sprintf("KEY_LOGOFF(0x%04x)", uint16(v))
	case KEY_DOLLAR:
		return fmt.Sprintf("KEY_DOLLAR(0x%04x)", uint16(v))
	case KEY_EURO:
		return fmt.Sprintf("KEY_EURO(0x%04x)", uint16(v))
	case KEY_FRAMEBACK:
		return fmt.Sprintf("KEY_FRAMEBACK(0x%04x)", uint16(v))
	case KEY_FRAMEFORWARD:
		return fmt.Sprintf("KEY_FRAMEFORWARD(0x%04x)", uint16(v))
	case KEY_CONTEXT_MENU:
		return fmt.Sprintf("KEY_CONTEXT_MENU(0x%04x)", uint16(v))
	case KEY_MEDIA_REPEAT:
		return fmt.Sprintf("KEY_MEDIA_REPEAT(0x%04x)", uint16(v))
	case KEY_10CHANNELSUP:
		return fmt.Sprintf("KEY_10CHANNELSUP(0x%04x)", uint16(v))
	case KEY_10CHANNELSDOWN:
		return fmt.Sprintf("KEY_10CHANNELSDOWN(0x%04x)", uint16(v))
	case KEY_IMAGES:
		return fmt.Sprintf("KEY_IMAGES(0x%04x)", uint16(v))
	case KEY_DEL_EOL:
		return fmt.Sprintf("KEY_DEL_EOL(0x%04x)", uint16(v))
	case KEY_DEL_EOS:
		return fmt.Sprintf("KEY_DEL_EOS(0x%04x)", uint16(v))
	case KEY_INS_LINE:
		return fmt.Sprintf("KEY_INS_LINE(0x%04x)", uint16(v))
	case KEY_DEL_LINE:
		return fmt.Sprintf("KEY_DEL_LINE(0x%04x)", uint16(v))
	case KEY_FN:
		return fmt.Sprintf("KEY_FN(0x%04x)", uint16(v))
	case KEY_FN_ESC:
		return fmt.Sprintf("KEY_FN_ESC(0x%04x)", uint16(v))
	case KEY_FN_F1:
		return fmt.Sprintf("KEY_FN_F1(0x%04x)", uint16(v))
	case KEY_FN_F2:
		return fmt.Sprintf("KEY_FN_F2(0x%04x)", uint16(v))
	case KEY_FN_F3:
		return fmt.Sprintf("KEY_FN_F3(0x%04x)", uint16(v))
	case KEY_FN_F4:
		return fmt.Sprintf("KEY_FN_F4(0x%04x)", uint16(v))
	case KEY_FN_F5:
		return fmt.Sprintf("KEY_FN_F5(0x%04x)", uint16(v))
	case KEY_FN_F6:
		return fmt.Sprintf("KEY_FN_F6(0x%04x)", uint16(v))
	case KEY_FN_F7:
		return fmt.Sprintf("KEY_FN_F7(0x%04x)", uint16(v))
	case KEY_FN_F8:
		return fmt.Sprintf("KEY_FN_F8(0x%04x)", uint16(v))
	case KEY_FN_F9:
		return fmt.Sprintf("KEY_FN_F9(0x%04x)", uint16(v))
	case KEY_FN_F10:
		return fmt.Sprintf("KEY_FN_F10(0x%04x)", uint16(v))
	case KEY_FN_F11:
		return fmt.Sprintf("KEY_FN_F11(0x%04x)", uint16(v))
	case KEY_FN_F12:
		return fmt.Sprintf("KEY_FN_F12(0x%04x)", uint16(v))
	case KEY_FN_1:
		return fmt.Sprintf("KEY_FN_1(0x%04x)", uint16(v))
	case KEY_FN_2:
		return fmt.Sprintf("KEY_FN_2(0x%04x)", uint16(v))
	case KEY_FN_D:
		return fmt.Sprintf("KEY_FN_D(0x%04x)", uint16(v))
	case KEY_FN_E:
		return fmt.Sprintf("KEY_FN_E(0x%04x)", uint16(v))
	case KEY_FN_F:
		return fmt.Sprintf("KEY_FN_F(0x%04x)", uint16(v))
	case KEY_FN_S:
		return fmt.Sprintf("KEY_FN_S(0x%04x)", uint16(v))
	case KEY_FN_B:
		return fmt.Sprintf("KEY_FN_B(0x%04x)", uint16(v))
	case KEY_BRL_DOT1:
		return fmt.Sprintf("KEY_BRL_DOT1(0x%04x)", uint16(v))
	case KEY_BRL_DOT2:
		return fmt.Sprintf("KEY_BRL_DOT2(0x%04x)", uint16(v))
	case KEY_BRL_DOT3:
		return fmt.Sprintf("KEY_BRL_DOT3(0x%04x)", uint16(v))
	case KEY_BRL_DOT4:
		return fmt.Sprintf("KEY_BRL_DOT4(0x%04x)", uint16(v))
	case KEY_BRL_DOT5:
		return fmt.Sprintf("KEY_BRL_DOT5(0x%04x)", uint16(v))
	case KEY_BRL_DOT6:
		return fmt.Sprintf("KEY_BRL_DOT6(0x%04x)", uint16(v))
	case KEY_BRL_DOT7:
		return fmt.Sprintf("KEY_BRL_DOT7(0x%04x)", uint16(v))
	case KEY_BRL_DOT8:
		return fmt.Sprintf("KEY_BRL_DOT8(0x%04x)", uint16(v))
	case KEY_BRL_DOT9:
		return fmt.Sprintf("KEY_BRL_DOT9(0x%04x)", uint16(v))
	case KEY_BRL_DOT10:
		return fmt.Sprintf("KEY_BRL_DOT10(0x%04x)", uint16(v))
	case KEY_NUMERIC_0:
		return fmt.Sprintf("KEY_NUMERIC_0(0x%04x)", uint16(v))
	case KEY_NUMERIC_1:
		return fmt.Sprintf("KEY_NUMERIC_1(0x%04x)", uint16(v))
	case KEY_NUMERIC_2:
		return fmt.Sprintf("KEY_NUMERIC_2(0x%04x)", uint16(v))
	case KEY_NUMERIC_3:
		return fmt.Sprintf("KEY_NUMERIC_3(0x%04x)", uint16(v))
	case KEY_NUMERIC_4:
		return fmt.Sprintf("KEY_NUMERIC_4(0x%04x)", uint16(v))
	case KEY_NUMERIC_5:
		return fmt.Sprintf("KEY_NUMERIC_5(0x%04x)", uint16(v))
	case KEY_NUMERIC_6:
		return fmt.Sprintf("KEY_NUMERIC_6(0x%04x)", uint16(v))
	case KEY_NUMERIC_7:
		return fmt.Sprintf("KEY_NUMERIC_7(0x%04x)", uint16(v))
	case KEY_NUMERIC_8:
		return fmt.Sprintf("KEY_NUMERIC_8(0x%04x)", uint16(v))
	case KEY_NUMERIC_9:
		return fmt.Sprintf("KEY_NUMERIC_9(0x%04x)", uint16(v))
	case KEY_NUMERIC_STAR:
		return fmt.Sprintf("KEY_NUMERIC_STAR(0x%04x)", uint16(v))
	case KEY_NUMERIC_POUND:
		return fmt.Sprintf("KEY_NUMERIC_POUND(0x%04x)", uint16(v))
	case KEY_NUMERIC_A:
		return fmt.Sprintf("KEY_NUMERIC_A(0x%04x)", uint16(v))
	case KEY_NUMERIC_B:
		return fmt.Sprintf("KEY_NUMERIC_B(0x%04x)", uint16(v))
	case KEY_NUMERIC_C:
		return fmt.Sprintf("KEY_NUMERIC_C(0x%04x)", uint16(v))
	case KEY_NUMERIC_D:
		return fmt.Sprintf("KEY_NUMERIC_D(0x%04x)", uint16(v))
	case KEY_CAMERA_FOCUS:
		return fmt.Sprintf("KEY_CAMERA_FOCUS(0x%04x)", uint16(v))
	case KEY_WPS_BUTTON:
		return fmt.Sprintf("KEY_WPS_BUTTON(0x%04x)", uint16(v))
	case KEY_TOUCHPAD_TOGGLE:
		return fmt.Sprintf("KEY_TOUCHPAD_TOGGLE(0x%04x)", uint16(v))
	case KEY_TOUCHPAD_ON:
		return fmt.Sprintf("KEY_TOUCHPAD_ON(0x%04x)", uint16(v))
	case KEY_TOUCHPAD_OFF:
		return fmt.Sprintf("KEY_TOUCHPAD_OFF(0x%04x)", uint16(v))
	case KEY_CAMERA_ZOOMIN:
		return fmt.Sprintf("KEY_CAMERA_ZOOMIN(0x%04x)", uint16(v))
	case KEY_CAMERA_ZOOMOUT:
		return fmt.Sprintf("KEY_CAMERA_ZOOMOUT(0x%04x)", uint16(v))
	case KEY_CAMERA_UP:
		return fmt.Sprintf("KEY_CAMERA_UP(0x%04x)", uint16(v))
	case KEY_CAMERA_DOWN:
		return fmt.Sprintf("KEY_CAMERA_DOWN(0x%04x)", uint16(v))
	case KEY_CAMERA_LEFT:
		return fmt.Sprintf("KEY_CAMERA_LEFT(0x%04x)", uint16(v))
	case KEY_CAMERA_RIGHT:
		return fmt.Sprintf("KEY_CAMERA_RIGHT(0x%04x)", uint16(v))
	case KEY_ATTENDANT_ON:
		return fmt.Sprintf("KEY_ATTENDANT_ON(0x%04x)", uint16(v))
	case KEY_ATTENDANT_OFF:
		return fmt.Sprintf("KEY_ATTENDANT_OFF(0x%04x)", uint16(v))
	case KEY_ATTENDANT_TOGGLE:
		return fmt.Sprintf("KEY_ATTENDANT_TOGGLE(0x%04x)", uint16(v))
	case KEY_LIGHTS_TOGGLE:
		return fmt.Sprintf("KEY_LIGHTS_TOGGLE(0x%04x)", uint16(v))
	case BTN_DPAD_UP:
		return fmt.Sprintf("BTN_DPAD_UP(0x%04x)", uint16(v))
	case BTN_DPAD_DOWN:
		return fmt.Sprintf("BTN_DPAD_DOWN(0x%04x)", uint16(v))
	case BTN_DPAD_LEFT:
		return fmt.Sprintf("BTN_DPAD_LEFT(0x%04x)", uint16(v))
	case BTN_DPAD_RIGHT:
		return fmt.Sprintf("BTN_DPAD_RIGHT(0x%04x)", uint16(v))
	case KEY_ALS_TOGGLE:
		return fmt.Sprintf("KEY_ALS_TOGGLE(0x%04x)", uint16(v))
	case KEY_BUTTONCONFIG:
		return fmt.Sprintf("KEY_BUTTONCONFIG(0x%04x)", uint16(v))
	case KEY_TASKMANAGER:
		return fmt.Sprintf("KEY_TASKMANAGER(0x%04x)", uint16(v))
	case KEY_JOURNAL:
		return fmt.Sprintf("KEY_JOURNAL(0x%04x)", uint16(v))
	case KEY_CONTROLPANEL:
		return fmt.Sprintf("KEY_CONTROLPANEL(0x%04x)", uint16(v))
	case KEY_APPSELECT:
		return fmt.Sprintf("KEY_APPSELECT(0x%04x)", uint16(v))
	case KEY_SCREENSAVER:
		return fmt.Sprintf("KEY_SCREENSAVER(0x%04x)", uint16(v))
	case KEY_VOICECOMMAND:
		return fmt.Sprintf("KEY_VOICECOMMAND(0x%04x)", uint16(v))
	case KEY_ASSISTANT:
		return fmt.Sprintf("KEY_ASSISTANT(0x%04x)", uint16(v))
	case KEY_BRIGHTNESS_MIN:
		return fmt.Sprintf("KEY_BRIGHTNESS_MIN(0x%04x)", uint16(v))
	case KEY_BRIGHTNESS_MAX:
		return fmt.Sprintf("KEY_BRIGHTNESS_MAX(0x%04x)", uint16(v))
	case KEY_KBDINPUTASSIST_PREV:
		return fmt.Sprintf("KEY_KBDINPUTASSIST_PREV(0x%04x)", uint16(v))
	case KEY_KBDINPUTASSIST_NEXT:
		return fmt.Sprintf("KEY_KBDINPUTASSIST_NEXT(0x%04x)", uint16(v))
	case KEY_KBDINPUTASSIST_PREVGROUP:
		return fmt.Sprintf("KEY_KBDINPUTASSIST_PREVGROUP(0x%04x)", uint16(v))
	case KEY_KBDINPUTASSIST_NEXTGROUP:
		return fmt.Sprintf("KEY_KBDINPUTASSIST_NEXTGROUP(0x%04x)", uint16(v))
	case KEY_KBDINPUTASSIST_ACCEPT:
		return fmt.Sprintf("KEY_KBDINPUTASSIST_ACCEPT(0x%04x)", uint16(v))
	case KEY_KBDINPUTASSIST_CANCEL:
		return fmt.Sprintf("KEY_KBDINPUTASSIST_CANCEL(0x%04x)", uint16(v))
	case KEY_RIGHT_UP:
		return fmt.Sprintf("KEY_RIGHT_UP(0x%04x)", uint16(v))
	case KEY_RIGHT_DOWN:
		return fmt.Sprintf("KEY_RIGHT_DOWN(0x%04x)", uint16(v))
	case KEY_LEFT_UP:
		return fmt.Sprintf("KEY_LEFT_UP(0x%04x)", uint16(v))
	case KEY_LEFT_DOWN:
		return fmt.Sprintf("KEY_LEFT_DOWN(0x%04x)", uint16(v))
	case KEY_ROOT_MENU:
		return fmt.Sprintf("KEY_ROOT_MENU(0x%04x)", uint16(v))
	case KEY_MEDIA_TOP_MENU:
		return fmt.Sprintf("KEY_MEDIA_TOP_MENU(0x%04x)", uint16(v))
	case KEY_NUMERIC_11:
		return fmt.Sprintf("KEY_NUMERIC_11(0x%04x)", uint16(v))
	case KEY_NUMERIC_12:
		return fmt.Sprintf("KEY_NUMERIC_12(0x%04x)", uint16(v))
	case KEY_AUDIO_DESC:
		return fmt.Sprintf("KEY_AUDIO_DESC(0x%04x)", uint16(v))
	case KEY_3D_MODE:
		return fmt.Sprintf("KEY_3D_MODE(0x%04x)", uint16(v))
	case KEY_NEXT_FAVORITE:
		return fmt.Sprintf("KEY_NEXT_FAVORITE(0x%04x)", uint16(v))
	case KEY_STOP_RECORD:
		return fmt.Sprintf("KEY_STOP_RECORD(0x%04x)", uint16(v))
	case KEY_PAUSE_RECORD:
		return fmt.Sprintf("KEY_PAUSE_RECORD(0x%04x)", uint16(v))
	case KEY_VOD:
		return fmt.Sprintf("KEY_VOD(0x%04x)", uint16(v))
	case KEY_UNMUTE:
		return fmt.Sprintf("KEY_UNMUTE(0x%04x)", uint16(v))
	case KEY_FASTREVERSE:
		return fmt.Sprintf("KEY_FASTREVERSE(0x%04x)", uint16(v))
	case KEY_SLOWREVERSE:
		return fmt.Sprintf("KEY_SLOWREVERSE(0x%04x)", uint16(v))
	case KEY_DATA:
		return fmt.Sprintf("KEY_DATA(0x%04x)", uint16(v))
	case KEY_ONSCREEN_KEYBOARD:
		return fmt.Sprintf("KEY_ONSCREEN_KEYBOARD(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY: // | BTN_TRIGGER_HAPPY1:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY|BTN_TRIGGER_HAPPY1(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY2:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY2(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY3:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY3(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY4:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY4(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY5:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY5(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY6:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY6(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY7:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY7(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY8:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY8(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY9:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY9(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY10:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY10(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY11:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY11(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY12:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY12(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY13:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY13(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY14:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY14(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY15:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY15(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY16:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY16(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY17:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY17(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY18:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY18(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY19:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY19(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY20:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY20(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY21:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY21(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY22:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY22(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY23:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY23(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY24:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY24(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY25:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY25(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY26:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY26(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY27:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY27(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY28:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY28(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY29:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY29(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY30:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY30(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY31:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY31(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY32:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY32(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY33:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY33(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY34:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY34(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY35:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY35(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY36:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY36(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY37:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY37(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY38:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY38(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY39:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY39(0x%04x)", uint16(v))
	case BTN_TRIGGER_HAPPY40:
		return fmt.Sprintf("BTN_TRIGGER_HAPPY40(0x%04x)", uint16(v))
	case KEY_MIN_INTERESTING: // KEY_MUTE:
		return fmt.Sprintf("KEY_MUTE|KEY_MIN_INTERESTING(0x%04x)", uint16(v))
	case KEY_MAX:
		return fmt.Sprintf("KEY_MAX(0x%04x)", uint16(v))
	case KEY_CNT:
		return fmt.Sprintf("KEY_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type REL_CODE uint16
type RelCode struct {
	v REL_CODE
}

func NewRelCode(v REL_CODE) *RelCode {
	return &RelCode{
		v: v,
	}
}

func (v *RelCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *RelCode) String() string {
	return v.v.String()
}

const (
	REL_X      REL_CODE = 0x00
	REL_Y               = 0x01
	REL_Z               = 0x02
	REL_RX              = 0x03
	REL_RY              = 0x04
	REL_RZ              = 0x05
	REL_HWHEEL          = 0x06
	REL_DIAL            = 0x07
	REL_WHEEL           = 0x08
	REL_MISC            = 0x09
	REL_MAX             = 0x0f
	REL_CNT             = (REL_MAX + 1)
)

func (v REL_CODE) String() string {
	switch v {
	case REL_X:
		return fmt.Sprintf("REL_X(0x%04x)", uint16(v))
	case REL_Y:
		return fmt.Sprintf("REL_Y(0x%04x)", uint16(v))
	case REL_Z:
		return fmt.Sprintf("REL_Z(0x%04x)", uint16(v))
	case REL_RX:
		return fmt.Sprintf("REL_RX(0x%04x)", uint16(v))
	case REL_RY:
		return fmt.Sprintf("REL_RY(0x%04x)", uint16(v))
	case REL_RZ:
		return fmt.Sprintf("REL_RZ(0x%04x)", uint16(v))
	case REL_HWHEEL:
		return fmt.Sprintf("REL_HWHEEL(0x%04x)", uint16(v))
	case REL_DIAL:
		return fmt.Sprintf("REL_DIAL(0x%04x)", uint16(v))
	case REL_WHEEL:
		return fmt.Sprintf("REL_WHEEL(0x%04x)", uint16(v))
	case REL_MISC:
		return fmt.Sprintf("REL_MISC(0x%04x)", uint16(v))
	case REL_MAX:
		return fmt.Sprintf("REL_MAX(0x%04x)", uint16(v))
	case REL_CNT:
		return fmt.Sprintf("REL_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type ABS_CODE uint16
type AbsCode struct {
	v ABS_CODE
}

func NewAbsCode(v ABS_CODE) *AbsCode {
	return &AbsCode{
		v: v,
	}
}

func (v *AbsCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *AbsCode) String() string {
	return v.v.String()
}

const (
	ABS_X              ABS_CODE = 0x00
	ABS_Y                       = 0x01
	ABS_Z                       = 0x02
	ABS_RX                      = 0x03
	ABS_RY                      = 0x04
	ABS_RZ                      = 0x05
	ABS_THROTTLE                = 0x06
	ABS_RUDDER                  = 0x07
	ABS_WHEEL                   = 0x08
	ABS_GAS                     = 0x09
	ABS_BRAKE                   = 0x0a
	ABS_HAT0X                   = 0x10
	ABS_HAT0Y                   = 0x11
	ABS_HAT1X                   = 0x12
	ABS_HAT1Y                   = 0x13
	ABS_HAT2X                   = 0x14
	ABS_HAT2Y                   = 0x15
	ABS_HAT3X                   = 0x16
	ABS_HAT3Y                   = 0x17
	ABS_PRESSURE                = 0x18
	ABS_DISTANCE                = 0x19
	ABS_TILT_X                  = 0x1a
	ABS_TILT_Y                  = 0x1b
	ABS_TOOL_WIDTH              = 0x1c
	ABS_VOLUME                  = 0x20
	ABS_MISC                    = 0x28
	ABS_MT_SLOT                 = 0x2f
	ABS_MT_TOUCH_MAJOR          = 0x30
	ABS_MT_TOUCH_MINOR          = 0x31
	ABS_MT_WIDTH_MAJOR          = 0x32
	ABS_MT_WIDTH_MINOR          = 0x33
	ABS_MT_ORIENTATION          = 0x34
	ABS_MT_POSITION_X           = 0x35
	ABS_MT_POSITION_Y           = 0x36
	ABS_MT_TOOL_TYPE            = 0x37
	ABS_MT_BLOB_ID              = 0x38
	ABS_MT_TRACKING_ID          = 0x39
	ABS_MT_PRESSURE             = 0x3a
	ABS_MT_DISTANCE             = 0x3b
	ABS_MT_TOOL_X               = 0x3c
	ABS_MT_TOOL_Y               = 0x3d
	ABS_MAX                     = 0x3f
	ABS_CNT                     = (ABS_MAX + 1)
)

func (v ABS_CODE) String() string {
	switch v {
	case ABS_X:
		return fmt.Sprintf("ABS_X(0x%04x)", uint16(v))
	case ABS_Y:
		return fmt.Sprintf("ABS_Y(0x%04x)", uint16(v))
	case ABS_Z:
		return fmt.Sprintf("ABS_Z(0x%04x)", uint16(v))
	case ABS_RX:
		return fmt.Sprintf("ABS_RX(0x%04x)", uint16(v))
	case ABS_RY:
		return fmt.Sprintf("ABS_RY(0x%04x)", uint16(v))
	case ABS_RZ:
		return fmt.Sprintf("ABS_RZ(0x%04x)", uint16(v))
	case ABS_THROTTLE:
		return fmt.Sprintf("ABS_THROTTLE(0x%04x)", uint16(v))
	case ABS_RUDDER:
		return fmt.Sprintf("ABS_RUDDER(0x%04x)", uint16(v))
	case ABS_WHEEL:
		return fmt.Sprintf("ABS_WHEEL(0x%04x)", uint16(v))
	case ABS_GAS:
		return fmt.Sprintf("ABS_GAS(0x%04x)", uint16(v))
	case ABS_BRAKE:
		return fmt.Sprintf("ABS_BRAKE(0x%04x)", uint16(v))
	case ABS_HAT0X:
		return fmt.Sprintf("ABS_HAT0X(0x%04x)", uint16(v))
	case ABS_HAT0Y:
		return fmt.Sprintf("ABS_HAT0Y(0x%04x)", uint16(v))
	case ABS_HAT1X:
		return fmt.Sprintf("ABS_HAT1X(0x%04x)", uint16(v))
	case ABS_HAT1Y:
		return fmt.Sprintf("ABS_HAT1Y(0x%04x)", uint16(v))
	case ABS_HAT2X:
		return fmt.Sprintf("ABS_HAT2X(0x%04x)", uint16(v))
	case ABS_HAT2Y:
		return fmt.Sprintf("ABS_HAT2Y(0x%04x)", uint16(v))
	case ABS_HAT3X:
		return fmt.Sprintf("ABS_HAT3X(0x%04x)", uint16(v))
	case ABS_HAT3Y:
		return fmt.Sprintf("ABS_HAT3Y(0x%04x)", uint16(v))
	case ABS_PRESSURE:
		return fmt.Sprintf("ABS_PRESSURE(0x%04x)", uint16(v))
	case ABS_DISTANCE:
		return fmt.Sprintf("ABS_DISTANCE(0x%04x)", uint16(v))
	case ABS_TILT_X:
		return fmt.Sprintf("ABS_TILT_X(0x%04x)", uint16(v))
	case ABS_TILT_Y:
		return fmt.Sprintf("ABS_TILT_Y(0x%04x)", uint16(v))
	case ABS_TOOL_WIDTH:
		return fmt.Sprintf("ABS_TOOL_WIDTH(0x%04x)", uint16(v))
	case ABS_VOLUME:
		return fmt.Sprintf("ABS_VOLUME(0x%04x)", uint16(v))
	case ABS_MISC:
		return fmt.Sprintf("ABS_MISC(0x%04x)", uint16(v))
	case ABS_MT_SLOT:
		return fmt.Sprintf("ABS_MT_SLOT(0x%04x)", uint16(v))
	case ABS_MT_TOUCH_MAJOR:
		return fmt.Sprintf("ABS_MT_TOUCH_MAJOR(0x%04x)", uint16(v))
	case ABS_MT_TOUCH_MINOR:
		return fmt.Sprintf("ABS_MT_TOUCH_MINOR(0x%04x)", uint16(v))
	case ABS_MT_WIDTH_MAJOR:
		return fmt.Sprintf("ABS_MT_WIDTH_MAJOR(0x%04x)", uint16(v))
	case ABS_MT_WIDTH_MINOR:
		return fmt.Sprintf("ABS_MT_WIDTH_MINOR(0x%04x)", uint16(v))
	case ABS_MT_ORIENTATION:
		return fmt.Sprintf("ABS_MT_ORIENTATION(0x%04x)", uint16(v))
	case ABS_MT_POSITION_X:
		return fmt.Sprintf("ABS_MT_POSITION_X(0x%04x)", uint16(v))
	case ABS_MT_POSITION_Y:
		return fmt.Sprintf("ABS_MT_POSITION_Y(0x%04x)", uint16(v))
	case ABS_MT_TOOL_TYPE:
		return fmt.Sprintf("ABS_MT_TOOL_TYPE(0x%04x)", uint16(v))
	case ABS_MT_BLOB_ID:
		return fmt.Sprintf("ABS_MT_BLOB_ID(0x%04x)", uint16(v))
	case ABS_MT_TRACKING_ID:
		return fmt.Sprintf("ABS_MT_TRACKING_ID(0x%04x)", uint16(v))
	case ABS_MT_PRESSURE:
		return fmt.Sprintf("ABS_MT_PRESSURE(0x%04x)", uint16(v))
	case ABS_MT_DISTANCE:
		return fmt.Sprintf("ABS_MT_DISTANCE(0x%04x)", uint16(v))
	case ABS_MT_TOOL_X:
		return fmt.Sprintf("ABS_MT_TOOL_X(0x%04x)", uint16(v))
	case ABS_MT_TOOL_Y:
		return fmt.Sprintf("ABS_MT_TOOL_Y(0x%04x)", uint16(v))
	case ABS_MAX:
		return fmt.Sprintf("ABS_MAX(0x%04x)", uint16(v))
	case ABS_CNT:
		return fmt.Sprintf("ABS_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type MSC_CODE uint16
type MscCode struct {
	v MSC_CODE
}

func NewMscCode(v MSC_CODE) *MscCode {
	return &MscCode{
		v: v,
	}
}

func (v *MscCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *MscCode) String() string {
	return v.v.String()
}

const (
	MSC_SERIAL    MSC_CODE = 0x00
	MSC_PULSELED           = 0x01
	MSC_GESTURE            = 0x02
	MSC_RAW                = 0x03
	MSC_SCAN               = 0x04
	MSC_TIMESTAMP          = 0x05
	MSC_MAX                = 0x07
	MSC_CNT                = (MSC_MAX + 1)
)

func (v MSC_CODE) String() string {
	switch v {
	case MSC_SERIAL:
		return fmt.Sprintf("MSC_SERIAL(0x%04x)", uint16(v))
	case MSC_PULSELED:
		return fmt.Sprintf("MSC_PULSELED(0x%04x)", uint16(v))
	case MSC_GESTURE:
		return fmt.Sprintf("MSC_GESTURE(0x%04x)", uint16(v))
	case MSC_RAW:
		return fmt.Sprintf("MSC_RAW(0x%04x)", uint16(v))
	case MSC_SCAN:
		return fmt.Sprintf("MSC_SCAN(0x%04x)", uint16(v))
	case MSC_TIMESTAMP:
		return fmt.Sprintf("MSC_TIMESTAMP(0x%04x)", uint16(v))
	case MSC_MAX:
		return fmt.Sprintf("MSC_MAX(0x%04x)", uint16(v))
	case MSC_CNT:
		return fmt.Sprintf("MSC_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type SW_CODE uint16
type SwCode struct {
	v SW_CODE
}

func NewSwCode(v SW_CODE) *SwCode {
	return &SwCode{
		v: v,
	}
}

func (v *SwCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *SwCode) String() string {
	return v.v.String()
}

const (
	SW_LID                  SW_CODE = 0x00
	SW_TABLET_MODE                  = 0x01
	SW_HEADPHONE_INSERT             = 0x02
	SW_RFKILL_ALL                   = 0x03
	SW_RADIO                        = SW_RFKILL_ALL
	SW_MICROPHONE_INSERT            = 0x04
	SW_DOCK                         = 0x05
	SW_LINEOUT_INSERT               = 0x06
	SW_JACK_PHYSICAL_INSERT         = 0x07
	SW_VIDEOOUT_INSERT              = 0x08
	SW_CAMERA_LENS_COVER            = 0x09
	SW_KEYPAD_SLIDE                 = 0x0a
	SW_FRONT_PROXIMITY              = 0x0b
	SW_ROTATE_LOCK                  = 0x0c
	SW_LINEIN_INSERT                = 0x0d
	SW_MUTE_DEVICE                  = 0x0e
	SW_PEN_INSERTED                 = 0x0f
	SW_MAX                          = 0x0f
	SW_CNT                          = (SW_MAX + 1)
)

func (v SW_CODE) String() string {
	switch v {
	case SW_LID:
		return fmt.Sprintf("SW_LID(0x%04x)", uint16(v))
	case SW_TABLET_MODE:
		return fmt.Sprintf("SW_TABLET_MODE(0x%04x)", uint16(v))
	case SW_HEADPHONE_INSERT:
		return fmt.Sprintf("SW_HEADPHONE_INSERT(0x%04x)", uint16(v))
	case SW_RFKILL_ALL: // | SW_RADIO:
		return fmt.Sprintf("SW_RFKILL_ALL|SW_RADIO(0x%04x)", uint16(v))
	case SW_MICROPHONE_INSERT:
		return fmt.Sprintf("SW_MICROPHONE_INSERT(0x%04x)", uint16(v))
	case SW_DOCK:
		return fmt.Sprintf("SW_DOCK(0x%04x)", uint16(v))
	case SW_LINEOUT_INSERT:
		return fmt.Sprintf("SW_LINEOUT_INSERT(0x%04x)", uint16(v))
	case SW_JACK_PHYSICAL_INSERT:
		return fmt.Sprintf("SW_JACK_PHYSICAL_INSERT(0x%04x)", uint16(v))
	case SW_VIDEOOUT_INSERT:
		return fmt.Sprintf("SW_VIDEOOUT_INSERT(0x%04x)", uint16(v))
	case SW_CAMERA_LENS_COVER:
		return fmt.Sprintf("SW_CAMERA_LENS_COVER(0x%04x)", uint16(v))
	case SW_KEYPAD_SLIDE:
		return fmt.Sprintf("SW_KEYPAD_SLIDE(0x%04x)", uint16(v))
	case SW_FRONT_PROXIMITY:
		return fmt.Sprintf("SW_FRONT_PROXIMITY(0x%04x)", uint16(v))
	case SW_ROTATE_LOCK:
		return fmt.Sprintf("SW_ROTATE_LOCK(0x%04x)", uint16(v))
	case SW_LINEIN_INSERT:
		return fmt.Sprintf("SW_LINEIN_INSERT(0x%04x)", uint16(v))
	case SW_MUTE_DEVICE:
		return fmt.Sprintf("SW_MUTE_DEVICE(0x%04x)", uint16(v))
	case SW_PEN_INSERTED: // | SW_MAX:
		return fmt.Sprintf("SW_PEN_INSERTED|SW_MAX(0x%04x)", uint16(v))
	case SW_CNT:
		return fmt.Sprintf("SW_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type LED_CODE uint16
type LedCode struct {
	v LED_CODE
}

func NewLedCode(v LED_CODE) *LedCode {
	return &LedCode{
		v: v,
	}
}

func (v *LedCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *LedCode) String() string {
	return v.v.String()
}

const (
	LED_NUML     LED_CODE = 0x00
	LED_CAPSL             = 0x01
	LED_SCROLLL           = 0x02
	LED_COMPOSE           = 0x03
	LED_KANA              = 0x04
	LED_SLEEP             = 0x05
	LED_SUSPEND           = 0x06
	LED_MUTE              = 0x07
	LED_MISC              = 0x08
	LED_MAIL              = 0x09
	LED_CHARGING          = 0x0a
	LED_MAX               = 0x0f
	LED_CNT               = (LED_MAX + 1)
)

func (v LED_CODE) String() string {
	switch v {
	case LED_NUML:
		return fmt.Sprintf("LED_NUML(0x%04x)", uint16(v))
	case LED_CAPSL:
		return fmt.Sprintf("LED_CAPSL(0x%04x)", uint16(v))
	case LED_SCROLLL:
		return fmt.Sprintf("LED_SCROLLL(0x%04x)", uint16(v))
	case LED_COMPOSE:
		return fmt.Sprintf("LED_COMPOSE(0x%04x)", uint16(v))
	case LED_KANA:
		return fmt.Sprintf("LED_KANA(0x%04x)", uint16(v))
	case LED_SLEEP:
		return fmt.Sprintf("LED_SLEEP(0x%04x)", uint16(v))
	case LED_SUSPEND:
		return fmt.Sprintf("LED_SUSPEND(0x%04x)", uint16(v))
	case LED_MUTE:
		return fmt.Sprintf("LED_MUTE(0x%04x)", uint16(v))
	case LED_MISC:
		return fmt.Sprintf("LED_MISC(0x%04x)", uint16(v))
	case LED_MAIL:
		return fmt.Sprintf("LED_MAIL(0x%04x)", uint16(v))
	case LED_CHARGING:
		return fmt.Sprintf("LED_CHARGING(0x%04x)", uint16(v))
	case LED_MAX:
		return fmt.Sprintf("LED_MAX(0x%04x)", uint16(v))
	case LED_CNT:
		return fmt.Sprintf("LED_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type REP_CODE uint16
type RepCode struct {
	v REP_CODE
}

func NewRepCode(v REP_CODE) *RepCode {
	return &RepCode{
		v: v,
	}
}

func (v *RepCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *RepCode) String() string {
	return v.v.String()
}

const (
	REP_DELAY  REP_CODE = 0x00
	REP_PERIOD          = 0x01
	REP_MAX             = 0x01
	REP_CNT             = (REP_MAX + 1)
)

func (v REP_CODE) String() string {
	switch v {
	case REP_DELAY:
		return fmt.Sprintf("REP_DELAY(0x%04x)", uint16(v))
	case REP_PERIOD: // | REP_MAX:
		return fmt.Sprintf("REP_PERIOD|REP_MAX(0x%04x)", uint16(v))
	case REP_CNT:
		return fmt.Sprintf("REP_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}

type SND_CODE uint16
type SndCode struct {
	v SND_CODE
}

func NewSndCode(v SND_CODE) *SndCode {
	return &SndCode{
		v: v,
	}
}

func (v *SndCode) ValueUint16() uint16 {
	return uint16(v.v)
}

func (v *SndCode) String() string {
	return v.v.String()
}

const (
	SND_CLICK SND_CODE = 0x00
	SND_BELL           = 0x01
	SND_TONE           = 0x02
	SND_MAX            = 0x07
	SND_CNT            = (SND_MAX + 1)
)

func (v SND_CODE) String() string {
	switch v {
	case SND_CLICK:
		return fmt.Sprintf("SND_CLICK(0x%04x)", uint16(v))
	case SND_BELL:
		return fmt.Sprintf("SND_BELL(0x%04x)", uint16(v))
	case SND_TONE:
		return fmt.Sprintf("SND_TONE(0x%04x)", uint16(v))
	case SND_MAX:
		return fmt.Sprintf("SND_MAX(0x%04x)", uint16(v))
	case SND_CNT:
		return fmt.Sprintf("SND_CNT(0x%04x)", uint16(v))
	}

	return fmt.Sprintf("UNKNOWN(0x%04x)", uint16(v))
}
