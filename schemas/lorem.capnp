using Go = import "/go.capnp";

@0xb1f92ff2dce1c0d3;

$Go.import("/go.capnp");
$Go.package("schemas");

struct LoremCpDataMap {
    index @0 :Int32;
    text @1 :Text;
}

struct LoremCp {
    id @0 :Text;
    data @1 :List(LoremCpDataMap);
    timestamp @2 :Int64;
}

