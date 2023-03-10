TESTS="$(make -s -C testbed list-stability-tests | xargs echo|sed 's/ /|/g')"

TESTS=(${TESTS//|/ })
MATRIX="{\"include\":["
curr=""
for i in "${!TESTS[@]}"; do
    curr="${TESTS[$i]}"
    MATRIX+="{\"test\":\"$curr\"},"
done
MATRIX+="]}"
echo "stabilitytest_matrix=$MATRIX" >> $GITHUB_OUTPUT
