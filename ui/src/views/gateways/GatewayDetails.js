import moment from "moment";
import React, { Component } from "react";
import { Line } from "react-chartjs-2";
import { Map, Marker } from 'react-leaflet';
import { Card, CardBody, CardTitle, Col, Row } from 'reactstrap';
import MapTileLayer from "../../components/MapTileLayer";
import i18n, { packageNS } from '../../i18n';
import GatewayStore from "../../stores/GatewayStore";





class GatewayDetails extends Component {
  constructor() {
    super();
    this.state = {};
    this.loadStats = this.loadStats.bind(this);
  }

  componentDidMount() {
    this.loadStats();
  }

  loadStats() {
    
    const end = moment().toISOString()
    const start = moment().subtract(30, "days").toISOString()

    GatewayStore.getStats(this.props.match.params.gatewayID, start, end, resp => {
      let statsDown = {
        labels: [],
        datasets: [
          {
            label: "rx received",
            borderColor: "rgba(33, 150, 243, 1)",
            backgroundColor: "rgba(0, 0, 0, 0)",
            lineTension: 0,
            pointBackgroundColor: "rgba(33, 150, 243, 1)",
            data: [],
          },
        ],
      }

      let statsUp = {
        labels: [],
        datasets: [
          {
            label: "tx emitted",
            borderColor: "rgba(33, 150, 243, 1)",
            backgroundColor: "rgba(0, 0, 0, 0)",
            lineTension: 0,
            pointBackgroundColor: "rgba(33, 150, 243, 1)",
            data: [],
          },
        ],
      }

      for (const row of resp.result) {
        statsUp.labels.push(moment(row.timestamp).format("Do"));
        statsDown.labels.push(moment(row.timestamp).format("Do"));
        statsUp.datasets[0].data.push(row.txPacketsEmitted);
        statsDown.datasets[0].data.push(row.rxPacketsReceivedOK);
      }

      this.setState({
        statsUp: statsUp,
        statsDown: statsDown,
      });
    });
  }

  render() {
    if (this.props.gateway === undefined || this.state.statsDown === undefined || this.state.statsUp === undefined) {
      return (<div></div>);
    }

    const style = {
      height: 322,
      zIndex: 1
    };

    const statsOptions = {
      legend: {
        display: false,
      },
      maintainAspectRatio: false,
      scales: {
        yAxes: [{
          ticks: {
            beginAtZero: true,
          },
        }],
      },
    }

    let position = [];
    if (typeof (this.props.gateway.location.latitude) !== "undefined" && typeof (this.props.gateway.location.longitude !== "undefined")) {
      position = [this.props.gateway.location.latitude, this.props.gateway.location.longitude];
    } else {
      position = [0, 0];
    }

    let lastseen = "";
    if (this.props.lastSeenAt !== undefined) {
      lastseen = moment(this.props.lastSeenAt).fromNow();
    }

    return (<React.Fragment>
      <Row>
        <Col lg={6}>
          <Card className="border shadow-none">
            <CardBody>
              <CardTitle tag="h4">{i18n.t(`${packageNS}:tr000423`)}</CardTitle>

              <h6 className="text-primary font-16">
                {i18n.t(`${packageNS}:tr000074`)}
              </h6>
              <p>
                {this.props.gateway.id}
              </p>

              <h6 className="text-primary font-16">
                {i18n.t(`${packageNS}:tr000432`)}
              </h6>
              <p>
                {this.props.gateway.location.altitude} meters
              </p>

              <h6 className="text-primary font-16">
                {i18n.t(`${packageNS}:tr000241`)}
              </h6>
              <p>
                {this.props.gateway.location.latitude}, {this.props.gateway.location.longitude}
              </p>

              <h6 className="text-primary font-16">
                {i18n.t(`${packageNS}:tr000242`)}
              </h6>
              <p>
                {lastseen}
              </p>
            </CardBody>
          </Card>
        </Col>

        <Col lg={6}>
          <Card className="border shadow-none">
            <CardBody className="p-1">
              <Map center={position} zoom={15} style={style} animate={true} scrollWheelZoom={false}>
                <MapTileLayer />
                <Marker position={position} />
              </Map>
            </CardBody>
          </Card>
        </Col>
      </Row>

      <Row>
        <Col lg={12}>
          <Card className="border shadow-none">
            <CardBody>
              <CardTitle tag="h4">{i18n.t(`${packageNS}:tr000243`)}</CardTitle>

              <div style={{ height: '300px' }}>
                <Line height={75} options={statsOptions} data={this.state.statsDown} redraw />
              </div>
            </CardBody>
          </Card>
        </Col>
      </Row>

      <Row>
        <Col lg={12}>
          <Card className="border shadow-none">
            <CardBody>
              <CardTitle tag="h4">{i18n.t(`${packageNS}:tr000244`)}</CardTitle>
              <div style={{ height: '300px' }}>
                <Line height={75} options={statsOptions} data={this.state.statsUp} redraw />
              </div>
            </CardBody>
          </Card>
        </Col>
      </Row>
    </React.Fragment>
    );
  }
}

export default GatewayDetails;
