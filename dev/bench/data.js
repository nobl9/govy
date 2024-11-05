window.BENCHMARK_DATA = {
  "lastUpdate": 1730809399757,
  "repoUrl": "https://github.com/nobl9/govy",
  "entries": {
    "Govy Go Benchmark": [
      {
        "commit": {
          "author": {
            "email": "48822818+nieomylnieja@users.noreply.github.com",
            "name": "Mateusz Hawrus",
            "username": "nieomylnieja"
          },
          "committer": {
            "email": "noreply@github.com",
            "name": "GitHub",
            "username": "web-flow"
          },
          "distinct": true,
          "id": "ddb940d6e370cacf2017a35db7057829a83fa4a5",
          "message": "chore: Split test workflows for pushes to main (#45)\n\n## Summary\r\n\r\nCurrently we're adding `if` statements to assert a GitHub event is equal\r\nto `push` to run some of the push-only actions.\r\nThis is not ideal as we're introducing misconfiguration potential and\r\nelevate the job's permissions.\r\nFurthermore, it seems that it might be easier to configure [benchmark\r\naction](https://github.com/benchmark-action/github-action-benchmark)\r\nseparately for publishing GitHub Pages and for checking the PR.",
          "timestamp": "2024-11-05T13:21:02+01:00",
          "tree_id": "5fc56448451f18b167fdde40b6c5821f2d58cb9e",
          "url": "https://github.com/nobl9/govy/commit/ddb940d6e370cacf2017a35db7057829a83fa4a5"
        },
        "date": 1730809399508,
        "tool": "go",
        "benches": [
          {
            "name": "BenchmarkEQ",
            "value": 672.1,
            "unit": "ns/op\t     224 B/op\t       6 allocs/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - ns/op",
            "value": 672.1,
            "unit": "ns/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - B/op",
            "value": 224,
            "unit": "B/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1791199 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ",
            "value": 778,
            "unit": "ns/op\t     240 B/op\t       6 allocs/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - ns/op",
            "value": 778,
            "unit": "ns/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - B/op",
            "value": 240,
            "unit": "B/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkNEQ - allocs/op",
            "value": 6,
            "unit": "allocs/op",
            "extra": "1541877 times\n4 procs"
          },
          {
            "name": "BenchmarkGT",
            "value": 814.1,
            "unit": "ns/op\t     400 B/op\t      10 allocs/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - ns/op",
            "value": 814.1,
            "unit": "ns/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - B/op",
            "value": 400,
            "unit": "B/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1472438 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE",
            "value": 743,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - ns/op",
            "value": 743,
            "unit": "ns/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkGTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1610506 times\n4 procs"
          },
          {
            "name": "BenchmarkLT",
            "value": 802.6,
            "unit": "ns/op\t     376 B/op\t      10 allocs/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - ns/op",
            "value": 802.6,
            "unit": "ns/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - B/op",
            "value": 376,
            "unit": "B/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLT - allocs/op",
            "value": 10,
            "unit": "allocs/op",
            "extra": "1407355 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE",
            "value": 767.7,
            "unit": "ns/op\t     368 B/op\t       8 allocs/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - ns/op",
            "value": 767.7,
            "unit": "ns/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - B/op",
            "value": 368,
            "unit": "B/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkLTE - allocs/op",
            "value": 8,
            "unit": "allocs/op",
            "extra": "1563001 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision",
            "value": 1197,
            "unit": "ns/op\t     520 B/op\t      18 allocs/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - ns/op",
            "value": 1197,
            "unit": "ns/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - B/op",
            "value": 520,
            "unit": "B/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkDurationPrecision - allocs/op",
            "value": 18,
            "unit": "allocs/op",
            "extra": "900819 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden",
            "value": 172.2,
            "unit": "ns/op\t     144 B/op\t       4 allocs/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - ns/op",
            "value": 172.2,
            "unit": "ns/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - B/op",
            "value": 144,
            "unit": "B/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkForbidden - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "7063042 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength",
            "value": 1317,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - ns/op",
            "value": 1317,
            "unit": "ns/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "848688 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength",
            "value": 1016,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - ns/op",
            "value": 1016,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength",
            "value": 1018,
            "unit": "ns/op\t     480 B/op\t      12 allocs/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - ns/op",
            "value": 1018,
            "unit": "ns/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - B/op",
            "value": 480,
            "unit": "B/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "991920 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength",
            "value": 1234,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - ns/op",
            "value": 1234,
            "unit": "ns/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceLength - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "888162 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength",
            "value": 1025,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - ns/op",
            "value": 1025,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength",
            "value": 1051,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - ns/op",
            "value": 1051,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength",
            "value": 1054,
            "unit": "ns/op\t     560 B/op\t      14 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - ns/op",
            "value": 1054,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - B/op",
            "value": 560,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapLength - allocs/op",
            "value": 14,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength",
            "value": 1025,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - ns/op",
            "value": 1025,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMinLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength",
            "value": 1079,
            "unit": "ns/op\t     544 B/op\t      12 allocs/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - ns/op",
            "value": 1079,
            "unit": "ns/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - B/op",
            "value": 544,
            "unit": "B/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkMapMaxLength - allocs/op",
            "value": 12,
            "unit": "allocs/op",
            "extra": "983133 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf",
            "value": 1105,
            "unit": "ns/op\t     536 B/op\t      22 allocs/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - ns/op",
            "value": 1105,
            "unit": "ns/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - B/op",
            "value": 536,
            "unit": "B/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOf - allocs/op",
            "value": 22,
            "unit": "allocs/op",
            "extra": "958146 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive",
            "value": 7148,
            "unit": "ns/op\t    3168 B/op\t      98 allocs/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - ns/op",
            "value": 7148,
            "unit": "ns/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - B/op",
            "value": 3168,
            "unit": "B/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkMutuallyExclusive - allocs/op",
            "value": 98,
            "unit": "allocs/op",
            "extra": "165039 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties",
            "value": 2614,
            "unit": "ns/op\t    1064 B/op\t      32 allocs/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - ns/op",
            "value": 2614,
            "unit": "ns/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - B/op",
            "value": 1064,
            "unit": "B/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkOneOfProperties - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "437707 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired",
            "value": 1052,
            "unit": "ns/op\t     704 B/op\t      23 allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - ns/op",
            "value": 1052,
            "unit": "ns/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - B/op",
            "value": 704,
            "unit": "B/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkRequired - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "1000000 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty",
            "value": 180.2,
            "unit": "ns/op\t     128 B/op\t       4 allocs/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - ns/op",
            "value": 180.2,
            "unit": "ns/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - B/op",
            "value": 128,
            "unit": "B/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringNotEmpty - allocs/op",
            "value": 4,
            "unit": "allocs/op",
            "extra": "6321806 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp",
            "value": 1492,
            "unit": "ns/op\t     644 B/op\t      20 allocs/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - ns/op",
            "value": 1492,
            "unit": "ns/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - B/op",
            "value": 644,
            "unit": "B/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "737235 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp",
            "value": 1496,
            "unit": "ns/op\t     693 B/op\t      20 allocs/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - ns/op",
            "value": 1496,
            "unit": "ns/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - B/op",
            "value": 693,
            "unit": "B/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDenyRegexp - allocs/op",
            "value": 20,
            "unit": "allocs/op",
            "extra": "742516 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel",
            "value": 15135,
            "unit": "ns/op\t    5640 B/op\t     154 allocs/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - ns/op",
            "value": 15135,
            "unit": "ns/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - B/op",
            "value": 5640,
            "unit": "B/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDNSLabel - allocs/op",
            "value": 154,
            "unit": "allocs/op",
            "extra": "77636 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII",
            "value": 4363,
            "unit": "ns/op\t    1552 B/op\t      41 allocs/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - ns/op",
            "value": 4363,
            "unit": "ns/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - B/op",
            "value": 1552,
            "unit": "B/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringASCII - allocs/op",
            "value": 41,
            "unit": "allocs/op",
            "extra": "267367 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID",
            "value": 16127,
            "unit": "ns/op\t   14085 B/op\t     217 allocs/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - ns/op",
            "value": 16127,
            "unit": "ns/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - B/op",
            "value": 14085,
            "unit": "B/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringUUID - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74044 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail",
            "value": 8593,
            "unit": "ns/op\t    3408 B/op\t     138 allocs/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - ns/op",
            "value": 8593,
            "unit": "ns/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - B/op",
            "value": 3408,
            "unit": "B/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEmail - allocs/op",
            "value": 138,
            "unit": "allocs/op",
            "extra": "137911 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL",
            "value": 8849,
            "unit": "ns/op\t    5904 B/op\t      54 allocs/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - ns/op",
            "value": 8849,
            "unit": "ns/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - B/op",
            "value": 5904,
            "unit": "B/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringURL - allocs/op",
            "value": 54,
            "unit": "allocs/op",
            "extra": "134198 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC",
            "value": 1057,
            "unit": "ns/op\t     752 B/op\t      25 allocs/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - ns/op",
            "value": 1057,
            "unit": "ns/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMAC - allocs/op",
            "value": 25,
            "unit": "allocs/op",
            "extra": "987164 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP",
            "value": 1599,
            "unit": "ns/op\t     824 B/op\t      31 allocs/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - ns/op",
            "value": 1599,
            "unit": "ns/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - B/op",
            "value": 824,
            "unit": "B/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIP - allocs/op",
            "value": 31,
            "unit": "allocs/op",
            "extra": "686692 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4",
            "value": 1700,
            "unit": "ns/op\t     896 B/op\t      32 allocs/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - ns/op",
            "value": 1700,
            "unit": "ns/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - B/op",
            "value": 896,
            "unit": "B/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv4 - allocs/op",
            "value": 32,
            "unit": "allocs/op",
            "extra": "662527 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6",
            "value": 1826,
            "unit": "ns/op\t    1056 B/op\t      36 allocs/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - ns/op",
            "value": 1826,
            "unit": "ns/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - B/op",
            "value": 1056,
            "unit": "B/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringIPv6 - allocs/op",
            "value": 36,
            "unit": "allocs/op",
            "extra": "554678 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR",
            "value": 3003,
            "unit": "ns/op\t    1576 B/op\t      63 allocs/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - ns/op",
            "value": 3003,
            "unit": "ns/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDR - allocs/op",
            "value": 63,
            "unit": "allocs/op",
            "extra": "380036 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4",
            "value": 5599,
            "unit": "ns/op\t    3048 B/op\t     118 allocs/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - ns/op",
            "value": 5599,
            "unit": "ns/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - B/op",
            "value": 3048,
            "unit": "B/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv4 - allocs/op",
            "value": 118,
            "unit": "allocs/op",
            "extra": "207870 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6",
            "value": 3544,
            "unit": "ns/op\t    2056 B/op\t      75 allocs/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - ns/op",
            "value": 3544,
            "unit": "ns/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - B/op",
            "value": 2056,
            "unit": "B/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCIDRv6 - allocs/op",
            "value": 75,
            "unit": "allocs/op",
            "extra": "328495 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON",
            "value": 1173,
            "unit": "ns/op\t     616 B/op\t      23 allocs/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - ns/op",
            "value": 1173,
            "unit": "ns/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - B/op",
            "value": 616,
            "unit": "B/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringJSON - allocs/op",
            "value": 23,
            "unit": "allocs/op",
            "extra": "927267 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains",
            "value": 2338,
            "unit": "ns/op\t    1448 B/op\t      44 allocs/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - ns/op",
            "value": 2338,
            "unit": "ns/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - B/op",
            "value": 1448,
            "unit": "B/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringContains - allocs/op",
            "value": 44,
            "unit": "allocs/op",
            "extra": "487446 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes",
            "value": 2383,
            "unit": "ns/op\t    1576 B/op\t      46 allocs/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - ns/op",
            "value": 2383,
            "unit": "ns/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - B/op",
            "value": 1576,
            "unit": "B/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringExcludes - allocs/op",
            "value": 46,
            "unit": "allocs/op",
            "extra": "475176 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith",
            "value": 1302,
            "unit": "ns/op\t     784 B/op\t      24 allocs/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - ns/op",
            "value": 1302,
            "unit": "ns/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - B/op",
            "value": 784,
            "unit": "B/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringStartsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "828422 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith",
            "value": 1298,
            "unit": "ns/op\t     752 B/op\t      24 allocs/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - ns/op",
            "value": 1298,
            "unit": "ns/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - B/op",
            "value": 752,
            "unit": "B/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringEndsWith - allocs/op",
            "value": 24,
            "unit": "allocs/op",
            "extra": "829464 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle",
            "value": 1598,
            "unit": "ns/op\t     976 B/op\t      30 allocs/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - ns/op",
            "value": 1598,
            "unit": "ns/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - B/op",
            "value": 976,
            "unit": "B/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringTitle - allocs/op",
            "value": 30,
            "unit": "allocs/op",
            "extra": "693373 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef",
            "value": 11688,
            "unit": "ns/op\t    3984 B/op\t      72 allocs/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - ns/op",
            "value": 11688,
            "unit": "ns/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - B/op",
            "value": 3984,
            "unit": "B/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringGitRef - allocs/op",
            "value": 72,
            "unit": "allocs/op",
            "extra": "101740 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath",
            "value": 37298,
            "unit": "ns/op\t    7577 B/op\t      99 allocs/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - ns/op",
            "value": 37298,
            "unit": "ns/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - B/op",
            "value": 7577,
            "unit": "B/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFileSystemPath - allocs/op",
            "value": 99,
            "unit": "allocs/op",
            "extra": "32180 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath",
            "value": 37738,
            "unit": "ns/op\t    7834 B/op\t     108 allocs/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - ns/op",
            "value": 37738,
            "unit": "ns/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - B/op",
            "value": 7834,
            "unit": "B/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringFilePath - allocs/op",
            "value": 108,
            "unit": "allocs/op",
            "extra": "31725 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath",
            "value": 37333,
            "unit": "ns/op\t    7449 B/op\t     103 allocs/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - ns/op",
            "value": 37333,
            "unit": "ns/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - B/op",
            "value": 7449,
            "unit": "B/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringDirPath - allocs/op",
            "value": 103,
            "unit": "allocs/op",
            "extra": "32030 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath",
            "value": 16020,
            "unit": "ns/op\t    8721 B/op\t     217 allocs/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - ns/op",
            "value": 16020,
            "unit": "ns/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - B/op",
            "value": 8721,
            "unit": "B/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringMatchFileSystemPath - allocs/op",
            "value": 217,
            "unit": "allocs/op",
            "extra": "74487 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp",
            "value": 156399419,
            "unit": "ns/op\t334255565 B/op\t  281343 allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - ns/op",
            "value": 156399419,
            "unit": "ns/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - B/op",
            "value": 334255565,
            "unit": "B/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringRegexp - allocs/op",
            "value": 281343,
            "unit": "allocs/op",
            "extra": "7 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab",
            "value": 46804,
            "unit": "ns/op\t   29417 B/op\t     614 allocs/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - ns/op",
            "value": 46804,
            "unit": "ns/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - B/op",
            "value": 29417,
            "unit": "B/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkStringCrontab - allocs/op",
            "value": 614,
            "unit": "allocs/op",
            "extra": "25482 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique",
            "value": 4194,
            "unit": "ns/op\t    2054 B/op\t      58 allocs/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - ns/op",
            "value": 4194,
            "unit": "ns/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - B/op",
            "value": 2054,
            "unit": "B/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkSliceUnique - allocs/op",
            "value": 58,
            "unit": "allocs/op",
            "extra": "277071 times\n4 procs"
          },
          {
            "name": "BenchmarkURL",
            "value": 1307,
            "unit": "ns/op\t     640 B/op\t      16 allocs/op",
            "extra": "854250 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - ns/op",
            "value": 1307,
            "unit": "ns/op",
            "extra": "854250 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - B/op",
            "value": 640,
            "unit": "B/op",
            "extra": "854250 times\n4 procs"
          },
          {
            "name": "BenchmarkURL - allocs/op",
            "value": 16,
            "unit": "allocs/op",
            "extra": "854250 times\n4 procs"
          }
        ]
      }
    ]
  }
}