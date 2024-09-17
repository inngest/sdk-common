make gen

if ! git diff --quiet; then
    echo "Generated code is out of date"
    exit 1
fi
