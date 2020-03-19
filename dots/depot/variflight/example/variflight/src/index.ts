import {grpc} from "@improbable-eng/grpc-web";
import {VariFlightDataService} from  "../../../typescript/protobuf/variflight_pb_service"
import {VariFlightData, GetFlightDataByFlightNumberRequest} from  "../../../typescript/protobuf/variflight_pb"

declare const USE_TLS: boolean
const host = USE_TLS ? "https://localhost:9090" : "http://localhost:9090"

// Testing query parameters
let forFlightNumberRequest = [
    {
        flightNumber: "CA4506",
        date: "2020-03-19",
    },
    {
        flightNumber: "CA4505",
        date: "2020-03-19",
    }
]

// Fake datum that may be responded.
let datum = [
    {
        FlightNo: "CA4506",
        FlightDeptimePlanDate: "2020-03-19",
        FlightDep: "南京",
        FlightArr: "成都",
        FlightDepAirport: "南京禄口",
        FlightArrAirport: "成都双流",
    },
    {
        FlightNo: "CA4505",
        FlightDeptimePlanDate: "2020-03-19",
        FlightDep: "南京",
        FlightArr: "成都",
        FlightDepAirport: "南京禄口",
        FlightArrAirport: "成都双流",
    },
    {
        FlightNo: "CA4504",
        FlightDeptimePlanDate: "2020-03-19",
        FlightDep: "南京",
        FlightArr: "成都",
        FlightDepAirport: "南京禄口",
        FlightArrAirport: "成都双流",
    },
    {
        FlightNo: "CA4503",
        FlightDeptimePlanDate: "2020-03-20",
        FlightDep: "南京",
        FlightArr: "成都",
        FlightDepAirport: "南京禄口",
        FlightArrAirport: "成都双流",
    },
    {
        FlightNo: "CA4502",
        FlightDeptimePlanDate: "2020-03-20",
        FlightDep: "南京",
        FlightArr: "成都",
        FlightDepAirport: "南京禄口",
        FlightArrAirport: "成都双流",
    },
    {
        FlightNo: "CA4501",
        FlightDeptimePlanDate: "2020-03-20",
        FlightDep: "南京",
        FlightArr: "成都",
        FlightDepAirport: "南京禄口",
        FlightArrAirport: "成都双流",
    },
]

// ---------- getFlightDataByFlightNumberRequest()---------
function getFlightDataByFlightNumberRequest() {
    const getByFlightNumberRequest = new GetFlightDataByFlightNumberRequest();
    const client = grpc.client(VariFlightDataService.GetFlightDataByFlightNumber, {
        host: host,
        transport: grpc.WebsocketTransport(),
        debug: true
    });
    console.log(`USE_TLS:${USE_TLS}, host:${host}`)
    client.onHeaders((headers: grpc.Metadata) => {
        console.log("queryBooks.onHeaders", headers);
    });
    client.onMessage((message: VariFlightData) => {
        console.log("queryBooks.onMessage", message.toObject());
    });
    client.onEnd((code: grpc.Code, msg: string, trailers: grpc.Metadata) => {
        console.log("queryBooks.onEnd", code, msg, trailers);
    });
    client.start();
    for (let req of forFlightNumberRequest) {
        getByFlightNumberRequest.setFlightnumber(req.flightNumber)
        getByFlightNumberRequest.setDate(req.date)
        client.send(getByFlightNumberRequest);
    }
    // client.finishSend();
    // client.close()
}

getFlightDataByFlightNumberRequest();

