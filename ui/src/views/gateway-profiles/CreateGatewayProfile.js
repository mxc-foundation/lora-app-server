import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';

import { withStyles } from '@material-ui/core/styles';
import Modal from '../../components/Modal';
import { Button, Breadcrumb, BreadcrumbItem, Form, FormGroup, Label, Input, FormText, Container, Row, Col, Card, CardBody } from 'reactstrap';



/* import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import { CardContent } from "@material-ui/core";


import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import DialogContentText from '@material-ui/core/DialogContentText';
import DialogTitle from '@material-ui/core/DialogTitle'; */

import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";

import GatewayProfileForm from "./GatewayProfileForm";
import GatewayProfileStore from "../../stores/GatewayProfileStore";
import NetworkServerStore from "../../stores/NetworkServerStore";


const styles = {
  card: {
    overflow: "visible",
  },
};


class CreateGatewayProfile extends Component {
  constructor() {
    super();
    this.state = {
      nsDialog: false,
    };
    this.onSubmit = this.onSubmit.bind(this);
    this.closeDialog = this.closeDialog.bind(this);
  }

  componentDidMount() {
    NetworkServerStore.list(0, 0, 0, resp => {
      if (resp.totalCount === "0") {
        this.setState({
          nsDialog: true,
        });
      }
    });
  }

  closeDialog() {
    this.setState({
      nsDialog: false,
    });
  }

  onSubmit(gatewayProfile) {
    GatewayProfileStore.create(gatewayProfile, resp => {
      this.props.history.push("/gateway-profiles");
    });
  }

  render() {
    return (
      <Form>
        {this.state.nsDialog && <Modal
          title={""}
          left={"DISMISS"}
          right={"ADD"}
          context={i18n.t(`${packageNS}:tr000377`)}
          callback={this.deleteGatewayProfile} />}

        <TitleBar>
          <Breadcrumb>
            <BreadcrumbItem><Link to={`/gateway-profiles`}>{i18n.t(`${packageNS}:tr000046`)}</Link></BreadcrumbItem>
            <BreadcrumbItem active>{i18n.t(`${packageNS}:tr000277`)}</BreadcrumbItem>
          </Breadcrumb>
        </TitleBar>
        <Row>
          <Col>
            <Card>
              <CardBody>
                <GatewayProfileForm
                  submitLabel={i18n.t(`${packageNS}:tr000277`)}
                  onSubmit={this.onSubmit}
                />
              </CardBody>
            </Card>
          </Col>
        </Row>
      </Form>

    );
  }
}

export default withStyles(styles)(withRouter(CreateGatewayProfile));
