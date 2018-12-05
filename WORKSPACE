load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "io_bazel_rules_go",
    urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.15.0/rules_go-0.15.0.tar.gz"],
    sha256 = "56d946edecb9879aed8dff411eb7a901f687e242da4fa95c81ca08938dd23bb4",
)

http_archive(
    name = "bazel_gazelle",
    urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.14.0/bazel-gazelle-0.14.0.tar.gz"],
    sha256 = "c0a5739d12c6d05b6c1ad56f2200cb0b57c5a70e03ebd2f7b87ce88cabf09c7b",
)

git_repository(
    name = "io_bazel_rules_docker",
    remote = "https://github.com/bazelbuild/rules_docker.git",
    tag = "v0.3.0",
)

new_local_repository(
    name = "secp256k1_lib",
    path = "third_party/libsecp256k1",
    build_file = "secp256k1_lib.BUILD",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

gazelle_dependencies()

# Import Go dependencies.
go_repository(
    name = "com_github_ethereum_go_ethereum",
    commit = "225171a4bfcc16bd12a1906b1e0d43d0b18c353b",
    importpath = "github.com/ethereum/go-ethereum",
)

go_repository(
    name = "com_github_aws_aws_sdk_go",
    commit = "ca229c7730be9278527fbb287f5a26c91e328d86",
    importpath = "github.com/aws/aws-sdk-go",
)

go_repository(
    name = "com_github_btcsuite_btcd",
    commit = "f899737d7f2764dc13e4d01ff00108ec58f766a9",
    importpath = "github.com/btcsuite/btcd",
)

go_repository(
    name = "com_github_deckarep_golang_set",
    commit = "1d4478f51bed434f1dadf96dcd9b43aabac66795",
    importpath = "github.com/deckarep/golang-set",
)

go_repository(
    name = "com_github_fatih_structs",
    commit = "a720dfa8df582c51dee1b36feabb906bde1588bd",
    importpath = "github.com/fatih/structs",
)

go_repository(
    name = "com_github_go_ini_ini",
    commit = "358ee7663966325963d4e8b2e1fbd570c5195153",
    importpath = "github.com/go-ini/ini",
)

go_repository(
    name = "com_github_go_stack_stack",
    commit = "259ab82a6cad3992b4e21ff5cac294ccb06474bc",
    importpath = "github.com/go-stack/stack",
)

go_repository(
    name = "com_github_golang_snappy",
    commit = "2e65f85255dbc3072edf28d6b5b8efc472979f5a",
    importpath = "github.com/golang/snappy",
)

go_repository(
    name = "com_github_jmespath_go_jmespath",
    commit = "0b12d6b5",
    importpath = "github.com/jmespath/go-jmespath",
)

go_repository(
    name = "com_github_lib_pq",
    commit = "90697d60dd844d5ef6ff15135d0203f65d2f53b8",
    importpath = "github.com/lib/pq",
)

go_repository(
    name = "com_github_mitchellh_mapstructure",
    commit = "f15292f7a699fcc1a38a80977f80a046874ba8ac",
    importpath = "github.com/mitchellh/mapstructure",
)

go_repository(
    name = "com_github_rs_cors",
    commit = "3fb1b69b103a84de38a19c3c6ec073dd6caa4d3f",
    importpath = "github.com/rs/cors",
)

go_repository(
    name = "com_github_syndtr_goleveldb",
    commit = "ae2bd5eed72d46b28834ec3f60db3a3ebedd8dbd",
    importpath = "github.com/syndtr/goleveldb",
)

go_repository(
    name = "in_gopkg_dgrijalva_jwt_go_v3",
    commit = "06ea1031745cb8b3dab3f6a236daf2b0aa468b7e",
    importpath = "gopkg.in/dgrijalva/jwt-go.v3",
)

go_repository(
    name = "in_gopkg_getstream_stream_go2_v1",
    commit = "31fed549e0fd8f5f5f7a7fcea18b33f47d3aca27",
    importpath = "gopkg.in/GetStream/stream-go2.v1",
)

go_repository(
    name = "in_gopkg_karalabe_cookiejar_v2",
    commit = "8dcd6a7f4951f6ff3ee9cbb919a06d8925822e57",
    importpath = "gopkg.in/karalabe/cookiejar.v2",
)

go_repository(
    name = "in_gopkg_leisurelink_httpsig_v1",
    commit = "6ea498b1c82dfe8c96d45df6c5877a520ace0494",
    importpath = "gopkg.in/LeisureLink/httpsig.v1",
)

go_repository(
    name = "in_gopkg_natefinch_npipe_v2",
    commit = "c1b8fa8bdccecb0b8db834ee0b92fdbcfa606dd6",
    importpath = "gopkg.in/natefinch/npipe.v2",
)

go_repository(
    name = "org_golang_x_net",
    commit = "aaf60122140d3fcf75376d319f0554393160eb50",
    importpath = "golang.org/x/net",
)

go_repository(
    name = "com_github_aws_aws_lambda_go",
    commit = "e630af386f8aad6ec9fc899b951cf62d37cfe56f",
    importpath = "github.com/aws/aws-lambda-go",
)

go_repository(
    name = "com_github_jmoiron_sqlx",
    commit = "0dae4fefe7c0e190f7b5a78dac28a1c82cc8d849",
    importpath = "github.com/jmoiron/sqlx",
)

go_repository(
    name = "com_github_davecgh_go_spew",
    commit = "8991bc29aa16c548c550c7ff78260e27b9ab7c73",
    importpath = "github.com/davecgh/go-spew",
)

go_repository(
    name = "com_github_pmezard_go_difflib",
    commit = "792786c7400a136282c1664665ae0a8db921c6c2",
    importpath = "github.com/pmezard/go-difflib",
)

go_repository(
    name = "com_github_stretchr_testify",
    commit = "f35b8ab0b5a2cef36673838d662e249dd9c94686",
    importpath = "github.com/stretchr/testify",
)
