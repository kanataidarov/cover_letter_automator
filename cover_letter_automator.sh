if [ "$#" -ne 3 ]; then
    echo "Usage: $0 <input_file> <rc_replacement> <rp_replacement>"
    echo "Example: $0 hh_en.txt TheCompany \"Java Developer\""
    exit 1
fi

input_file="$1"
rc="$2"
rp="$3"

go run cmd/main/main.go -rc "$rc" -rp "$rp" "template/$input_file"