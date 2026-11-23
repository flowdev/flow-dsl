# Test Mermaid Links

```mermaid
flowchart LR
    user("UserOperation")
    user_out1["out1"]
    user_data1([data1])
    user_data2([data2])
    user_err["err"]
    user_error(["error"])
    dash("Dashboard")
    dash_in1["in1"]
    dash_err["err"]
    dash_error(["error"])
    err(("Error"))

    user --- user_out1
    user_out1 --- user_data1 --- user_data2 --> dash_in1
    dash_in1 --- dash
    dash --- dash_err
    dash_err --- dash_error --> err
    user --- user_err
    user_err --- user_error --> err

    click dash "https://google.com"
    click user "../README.md"
    classDef comp fill:#239,color:#ccc;
    classDef port fill:#239,color:#ccc,stroke-width:0;
    class user comp
    class user_out1 port
    class user_err port
    class dash_in1 port
    class dash comp
    class dash_err port
```