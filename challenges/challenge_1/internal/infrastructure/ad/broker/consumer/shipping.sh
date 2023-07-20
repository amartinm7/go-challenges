#!/bin/bash
kcat -v -b localhost:9092 -t topic.test.1 -P -T shipping-event.json
