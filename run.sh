EXICUTE_FILE="build/tw-bn-detector"

while getopts "rbafd:" OPTION; do
    case "$OPTION" in
        r)
            $EXICUTE_FILE
            ;;
        b)
            go build -o $EXICUTE_FILE main.go
            ;;
        a)
            go build -o $EXICUTE_FILE main.go
            $EXICUTE_FILE
            ;;
        f)
            go fmt ./...
            ;;
        d)
            EXICUTE_FILE="$OPTARG"
            ;;
    esac
done

if [ $# -eq 0 ]
  then
    go build -o $EXICUTE_FILE main.go
    $EXICUTE_FILE
fi