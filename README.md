# oilking
A trading bot using bitflyer api.

If you use the code, please take responsibility at your own risk. We do not take any responsibility even if you suffer a loss.

# Get started
## Set enviroment variables
Copy an `.env.example` as a `.env` and edit it.
`cp .env.example .env`

## Build
Build the source code.
`make docker-compose-build`

## Run
Run the binary file.
`make docker-compose-up` or `make docker-compose-up-d`

# Supporting API
## HTTP API
### HTTP Public API
|      Support       | Method |     Endpoint      |
| ------------------ | ------ | ----------------- |
|                    | GET    | /v1/getmarkets    |
|                    | GET    | /v1/markets       |
|                    | GET    | /v1/getboard      |
|                    | GET    | /v1/board         |
| :white_check_mark: | GET    | /v1/getticker     |
|                    | GET    | /v1/getexecutions |
| :white_check_mark: | GET    | /v1/getchats      |
|                    | GET    | /v1/gethealth     |
|                    | GET    | /v1/getboardstate |

### HTTP Private API
|      Support       | Method |           Endpoint           |
| ------------------ | ------ | ---------------------------- |
|                    | GET    | /v1/me/getpermissions        |
| :white_check_mark: | GET    | /v1/me/getbalance            |
| :white_check_mark: | GET    | /v1/me/getcollateral         |
|                    | GET    | /v1/me/getcollateralaccounts |
| :white_check_mark: | POST   | /v1/me/sendchildorder        |
|                    | POST   | /v1/me/sendparentorder       |
| :white_check_mark: | POST   | /v1/me/cancelchildorder      |
|                    | POST   | /v1/me/cancelparentorder     |
| :white_check_mark: | POST   | /v1/me/cancelallchildorders  |
| :white_check_mark: | GET    | /v1/me/getchildorders        |
|                    | GET    | /v1/me/getparentorders       |
|                    | GET    | /v1/me/getparentorder        |
| :white_check_mark: | GET    | /v1/me/getexecutions         |
| :white_check_mark: | GET    | /v1/me/getbalancehistory     |
| :white_check_mark: | GET    | /v1/me/getpositions          |
|                    | GET    | /v1/me/getcollateralhistory  |
| :white_check_mark: | GET    | /v1/me/gettradingcommission  |
|                    | GET    | /v1/me/getaddresses          |
|                    | GET    | /v1/me/getcoinins            |
|                    | GET    | /v1/me/getcoinouts           |
|                    | GET    | /v1/me/getdeposits           |
|                    | GET    | /v1/me/getwithdrawals        |
|                    | GET    | /v1/me/getbankaccounts       |
|                    | POST   | /v1/me/withdraw              |

## Realtime API（JSON-RPC 2.0 over WebSocket）
### Public Channels
|      Support       |              Channel name               |
| ------------------ | --------------------------------------- |
|                    | lightning_board_snapshot_{product_code} |
|                    | lightning_board_{product_code}          |
| :white_check_mark: | lightning_ticker_{product_code}         |
|                    | lightning_executions_{product_code}     |


### Private Channels
| Support |              Channel name               |
| ------- | --------------------------------------- |
|         | lightning_board_snapshot_{product_code} |
|         | parent_order_events                     |

# Reference
- [lightning.bitflyer/docs](https://lightning.bitflyer.com/docs)

# Contributing
We welcome your issue or pull request from everyone.

We'd appreciate if you could contribute our project.

Any request or question is OK.

日本語でも大丈夫です :D

# License
This project is licensed under the terms of the MIT license.

# Author
bmf - Software engineer.

- [github - bmf-san/bmf-san](https://github.com/bmf-san/bmf-san)
- [twitter - @bmf-san](https://twitter.com/bmf_san)
- [blog - bmf-tech](http://bmf-tech.com/)
