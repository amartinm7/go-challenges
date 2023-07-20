
# for adId in "87f407fe-2205-4ba9-951a-036e2e7c0ffa", "411355b1-8ecd-474b-ab6a-8f845875477e", "03ac1f1f-ec36-4dfe-94c6-7905f76a4ee8", "1f878d26-1eaf-47e1-92be-0f6a86d85a21", "82794288-79c6-4a8d-8c3f-9c7f0eb07336", "cd6ef1e3-6b0e-4c97-81e0-62d8a40b4488", "a65610e0-da91-4389-9831-9b51c2c2a373", "df8d4f8c-db0a-4266-87cf-22fa487a5eb1", "92a2c7ba-dd2f-4dbf-b3c9-8f6f51436a29", "210114b0-a80d-4fbe-8c49-6fd717f00f37", "211f26c4-78c0-4cc1-ad06-5ddb13807965", "c46c00b8-212e-4917-9250-18aec1fb7db3", "673c0afe-1b42-40c3-8156-12f5b04a4856", "2fa488a3-a218-49a9-8de6-229915f353eb", "52a7a635-c787-4596-bdfb-1081f82c9783", "09751359-b25b-48b2-b95a-1b79e9332bdb", "9dba9ef0-42c6-4e1e-9920-aed27d628357", "650b3dba-a591-43cc-b877-9bb7747277f1", "38869e76-9f70-47f8-a307-a56fbb3d0bb3", "7d07ac99-32de-4cdd-b280-8a229f8f44ae", "ced17635-9204-417e-a9a8-bdcf8af8cb01", "d15eee3e-64ed-4c5f-b771-e9325f8547b3", "36631cf2-128f-496d-94a1-01d2fccabbc7", "1bf8960c-2f03-4435-8f4a-4f1eefb5d85c", "9055411f-81f0-4852-8c09-c70beb7737a9", "65d3df81-35b0-4384-99b1-0651885c2205", "df8b7373-18da-41ea-9b17-2667c2d41d75", "049eded1-c2fd-4fa8-b038-77c76a36ebab", "0f5097dc-3ebc-4a70-8b21-ea7c67bd0ee8", "f78f0b25-8d64-4d88-9ce4-b0d54cede948"
for adId in "87f407fe-2205-4ba9-951a-036e2e7c0ffa" "411355b1-8ecd-474b-ab6a-8f845875477e" "03ac1f1f-ec36-4dfe-94c6-7905f76a4ee8" "1f878d26-1eaf-47e1-92be-0f6a86d85a21" "82794288-79c6-4a8d-8c3f-9c7f0eb07336" "cd6ef1e3-6b0e-4c97-81e0-62d8a40b4488" "a65610e0-da91-4389-9831-9b51c2c2a373" "df8d4f8c-db0a-4266-87cf-22fa487a5eb1" "92a2c7ba-dd2f-4dbf-b3c9-8f6f51436a29" "210114b0-a80d-4fbe-8c49-6fd717f00f37" "211f26c4-78c0-4cc1-ad06-5ddb13807965" "c46c00b8-212e-4917-9250-18aec1fb7db3" "673c0afe-1b42-40c3-8156-12f5b04a4856" "2fa488a3-a218-49a9-8de6-229915f353eb" "52a7a635-c787-4596-bdfb-1081f82c9783" "09751359-b25b-48b2-b95a-1b79e9332bdb" "9dba9ef0-42c6-4e1e-9920-aed27d628357" "650b3dba-a591-43cc-b877-9bb7747277f1" "38869e76-9f70-47f8-a307-a56fbb3d0bb3" "7d07ac99-32de-4cdd-b280-8a229f8f44ae" "ced17635-9204-417e-a9a8-bdcf8af8cb01" "d15eee3e-64ed-4c5f-b771-e9325f8547b3" "36631cf2-128f-496d-94a1-01d2fccabbc7" "1bf8960c-2f03-4435-8f4a-4f1eefb5d85c" "9055411f-81f0-4852-8c09-c70beb7737a9" "65d3df81-35b0-4384-99b1-0651885c2205" "df8b7373-18da-41ea-9b17-2667c2d41d75" "049eded1-c2fd-4fa8-b038-77c76a36ebab" "0f5097dc-3ebc-4a70-8b21-ea7c67bd0ee8" "f78f0b25-8d64-4d88-9ce4-b0d54cede948"
do
  # shellcheck disable=SC2059
  printf "Post ad: $adId\n"
  curl -X POST -v --location 'http://localhost:8000/v1/ad' \
  --header 'Content-Type: application/json' \
  --data "{
    \"id\": \"$adId\",
    \"title\": \"opel astra\",
    \"description\": \"como nuevo\",
    \"price\": 15000,
    \"timeStamp\": \"2022-12-01\"
  }"
  printf "\nmessage sent"
  sleep 1
  # shellcheck disable=SC2059
  printf "\nSubmmited adId: $adId\n"
done

