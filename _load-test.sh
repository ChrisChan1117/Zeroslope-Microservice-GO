#!/bin/bash
echo "POST http://localhost:8080/samples/" | vegeta attack -body tests/data/creates.json -duration=30s -rate=1000 | tee TESTRESULTS-CREATE.bin | vegeta report
echo "DELETE http://localhost:8080/samples/" | vegeta attack -body tests/data/deletes.json -duration=30s -rate=1000 | tee TESTRESULTS-DELETE.bin | vegeta report
