import React, { Component } from "react";
import { withRouter, Link } from 'react-router-dom';

import Admin from "../../components/Admin";
import { Breadcrumb, BreadcrumbItem, Button, Card, Container, Modal, ModalHeader, ModalBody, ModalFooter, NavLink, Row, Col } from 'reactstrap';
import { withStyles } from "@material-ui/core/styles";

import i18n, { packageNS } from '../../i18n';
import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import Loader from "../../components/Loader";

import ApplicationStore from "../../stores/ApplicationStore";
import DeviceProfileStore from "../../stores/DeviceProfileStore";
import DeviceStore from "../../stores/DeviceStore";
import DeviceForm from "./DeviceForm";

import OrganizationStore from "../../stores/OrganizationStore";

import breadcrumbStyles from "../common/BreadcrumbStyles";

const localStyles = {};

const styles = {
  ...breadcrumbStyles,
  ...localStyles
};


class CreateDevice extends Component {
  constructor() {
    super();
    this.state = {
      appDialog: false,
      dpDialog: false,
      loading: true,
    };
  }

  componentDidMount() {
    const currentOrgID = this.props.organizationID || this.props.match.params.organizationID;
    const currentApplicationID = this.props.match.params.applicationID;
    if (currentApplicationID === undefined) {
      this.setState({
        appDialog: true
      });
    } else {
      ApplicationStore.get(currentApplicationID, resp => {
        this.setState({
          application: resp,
        });
      });

      DeviceProfileStore.list(0, currentApplicationID, 0, 0, resp => {
        if (resp.totalCount === "0") {
          this.setState({
            dpDialog: true,
            loading: false
          });
        }
      });
    }

    OrganizationStore.get(currentOrgID, resp => {
      this.setState({
        organization: resp.organization,
        loading: false
      });
    });
  }

  toggleAppDialog = () => {
    this.setState({
      appDialog: !this.state.appDialog,
    });
  }

  toggleDpDialog = () => {
    this.setState({
      dpDialog: !this.state.dpDialog,
    });
  }

  onSubmit = (device) => {
    const currentOrgID = this.props.organizationID || this.props.match.params.organizationID;
    const deviceApplicationID = device.applicationID;
    const isDeviceApplication = deviceApplicationID !== undefined;

    DeviceStore.create(device, resp => {
      if (!isDeviceApplication) {
        this.props.history.push(`/organizations/${this.props.match.params.organizationID}/devices`);
      }

      DeviceProfileStore.get(device.deviceProfileID, resp => {
        if (resp.deviceProfile.supportsJoin) {
          isDeviceApplication
            ? this.props.history.push(`/organizations/${currentOrgID}/applications/${deviceApplicationID}/devices/${device.devEUI}/keys`)
            : this.props.history.push(`/organizations/${currentOrgID}/devices/${device.devEUI}/keys`);
        } else {
          isDeviceApplication
            ? this.props.history.push(`/organizations/${currentOrgID}/applications/${deviceApplicationID}/devices/${device.devEUI}/activation`)
            : this.props.history.push(`/organizations/${currentOrgID}/devices/${device.devEUI}/activation`);
        }
      });

    });
  }

  render() {
    const { application, device, loading, organization } = this.state;
    const { classes, match } = this.props;
    const currentOrgID = this.props.organizationID || this.props.match.params.organizationID;
    const currentApplicationID = this.props.applicationID || this.props.match.params.applicationID;
    const currentOrgName = organization && (organization.name || organization.displayName);
    const currentAppName = application && (application.application.name);

    const closeAppBtn = <button className="close" onClick={this.toggleAppDialog}>&times;</button>;
    const closeDpBtn = <button className="close" onClick={this.toggleDpDialog}>&times;</button>;

    return (
      <React.Fragment>
        <Modal
          isOpen={this.state.appDialog}
          toggle={this.toggleAppDialog}
          aria-labelledby="help-dialog-title"
          aria-describedby="help-dialog-description"
        >
          <ModalHeader
            toggle={this.toggleAppDialog}
            close={closeAppBtn}
            id="help-dialog-title"
          >
            Add an Application?
          </ModalHeader>
          <ModalBody id="help-dialog-description">
            <p>
              You can create an application for your device to belong to.
            </p>
            <p>
              Would you like to create an application?
            </p>
          </ModalBody>
          <ModalFooter>
            <Button variant="outlined">
              <NavLink
                style={{ color: "#fff", padding: "0" }}
                tag={Link}
                to={`/organizations/${currentOrgID}/applications/create`}
              >
                {i18n.t(`${packageNS}:tr000277`)}
              </NavLink>
            </Button>
            <Button color="primary" onClick={this.toggleAppDialog}>{i18n.t(`${packageNS}:tr000166`)}</Button>{' '}
          </ModalFooter>
        </Modal>

        <Modal
          isOpen={this.state.dpDialog}
          toggle={this.toggleDpDialog}
          aria-labelledby="help-dialog-title"
          aria-describedby="help-dialog-description"
        >
          <ModalHeader
            toggle={this.toggleDpDialog}
            close={closeDpBtn}
            id="help-dialog-title"
          >
            Add a Device Profile?
          </ModalHeader>
          <ModalBody id="help-dialog-description">
            <p>
              The selected application does not have access to any device-profiles.
              A device-profile defines the capabilities and boot parameters of a device. You can create multiple device-profiles for different kind of devices.
            </p>
            <p>
              Would you like to create a device-profile?
            </p>
          </ModalBody>
          <ModalFooter>
            <Button variant="outlined">
              <NavLink
                style={{ color: "#fff", padding: "0" }}
                tag={Link}
                to={`/organizations/${currentOrgID}/device-profiles/create`}
              >
                {i18n.t(`${packageNS}:tr000277`)}
              </NavLink>
            </Button>
            <Button color="primary" onClick={this.toggleDpDialog}>{i18n.t(`${packageNS}:tr000166`)}</Button>{' '}
          </ModalFooter>
        </Modal>

        <TitleBar>
          {
            currentApplicationID ? (
              <Breadcrumb className={classes.breadcrumb}>
                <Admin>
                  <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={
                    `/organizations/${currentOrgID}/applications`
                  }>{i18n.t(`${packageNS}:tr000076`)}</Link>
                  </BreadcrumbItem>
                  <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={
                    `/organizations/${currentOrgID}/applications/${currentApplicationID}`
                  }>{currentAppName || currentApplicationID || '???'}</Link>
                  </BreadcrumbItem>
                </Admin>
                <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={
                  `/organizations/${currentOrgID}/applications/${currentApplicationID}`
                }>{i18n.t(`${packageNS}:tr000278`)}</Link></BreadcrumbItem>
                <BreadcrumbItem active>{i18n.t(`${packageNS}:tr000277`)}</BreadcrumbItem>
              </Breadcrumb>
            ) : (
                <Breadcrumb className={classes.breadcrumb}>
                  <Admin>
                    <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={
                      `/organizations`
                    }>Organizations</Link></BreadcrumbItem>
                    <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={
                      `/organizations/${currentOrgID}`
                    }>{currentOrgName || currentOrgID}</Link></BreadcrumbItem>
                  </Admin>
                  <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={
                    `/organizations/${currentOrgID}/devices`
                  }>{i18n.t(`${packageNS}:tr000278`)}</Link></BreadcrumbItem>
                  <BreadcrumbItem active>{i18n.t(`${packageNS}:tr000277`)}</BreadcrumbItem>
                </Breadcrumb>
              )
          }
        </TitleBar>

        {/* <Card className="card-box shadow-sm" style={{ minWidth: "25rem" }}> */}
        {/* <Card body> */}
        <DeviceForm
          submitLabel={i18n.t(`${packageNS}:tr000277`)}
          onSubmit={this.onSubmit}
          match={this.props.match}
          loading={loading}
        />
        <br />
      </React.Fragment>
    );
  }
}

export default withStyles(styles)(withRouter(CreateDevice));
