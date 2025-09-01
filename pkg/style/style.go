package style

var (
    RESET = "\033[0m"

    // Text styles
    BOLD          = "\033[1m"
    DIM           = "\033[2m"
    ITALIC        = "\033[3m" // !!! May not work everywhere !!!
    UNDERLINE     = "\033[4m"
    BLINK         = "\033[5m" // !!! May not work everywhere !!!
    INVERT        = "\033[7m"
    STRIKETHROUGH = "\033[9m"
    HIDDEN        = "\033[8m"

    // Foreground colors
    FG_BLACK   = "\033[30m"
    FG_RED     = "\033[31m"
    FG_GREEN   = "\033[32m"
    FG_YELLOW  = "\033[33m"
    FG_BLUE    = "\033[34m"
    FG_MAGENTA = "\033[35m"
    FG_CYAN    = "\033[36m"
    FG_WHITE   = "\033[37m"

    // Bright foreground colors
    FG_BRIGHT_BLACK   = "\033[90m"
    FG_BRIGHT_RED     = "\033[91m"
    FG_BRIGHT_GREEN   = "\033[92m"
    FG_BRIGHT_YELLOW  = "\033[93m"
    FG_BRIGHT_BLUE    = "\033[94m"
    FG_BRIGHT_MAGENTA = "\033[95m"
    FG_BRIGHT_CYAN    = "\033[96m"
    FG_BRIGHT_WHITE   = "\033[97m"

    // Background colors
    BG_BLACK   = "\033[40m"
    BG_RED     = "\033[41m"
    BG_GREEN   = "\033[42m"
    BG_YELLOW  = "\033[43m"
    BG_BLUE    = "\033[44m"
    BG_MAGENTA = "\033[45m"
    BG_CYAN    = "\033[46m"
    BG_WHITE   = "\033[47m"

    // Bright background colors
    BG_BRIGHT_BLACK   = "\033[100m"
    BG_BRIGHT_RED     = "\033[101m"
    BG_BRIGHT_GREEN   = "\033[102m"
    BG_BRIGHT_YELLOW  = "\033[103m"
    BG_BRIGHT_BLUE    = "\033[104m"
    BG_BRIGHT_MAGENTA = "\033[105m"
    BG_BRIGHT_CYAN    = "\033[106m"
    BG_BRIGHT_WHITE   = "\033[107m"
)
