import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardContent from "@material-ui/core/CardContent";

import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import AESKeyField from "../../components/AESKeyField";
import DeviceStore from "../../stores/DeviceStore";


class LW11DeviceKeysForm extends FormComponent {
  render() {
    let object = {};
    if (this.props.object !== undefined) {
      object = this.props.object;
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.onSubmit}
      >
        <AESKeyField
          id="nwkKey"
          label={i18n.t(`${packageNS}:tr000385`)}
          helperText={i18n.t(`${packageNS}:tr000386`)}
          onChange={this.onChange}
          value={object.nwkKey || ""}
          margin="normal"
          fullWidth
          required
          random
        />
        <AESKeyField
          id="appKey"
          label={i18n.t(`${packageNS}:tr000387`)}
          helperText={i18n.t(`${packageNS}:tr000386`)}
          onChange={this.onChange}
          value={object.appKey || ""}
          margin="normal"
          fullWidth
          required
          random
        />
      </Form>
    );
  }
}

class LW10DeviceKeysForm extends FormComponent {
  render() {
    let object = {};
    if (this.props.object !== undefined) {
      object = this.props.object;
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.onSubmit}
      >
        <AESKeyField
          id="nwkKey"
          label={i18n.t(`${packageNS}:tr000388`)}
          helperText="For LoRaWAN 1.0 devices. In case your device supports LoRaWAN 1.1, update the device-profile first."
          onChange={this.onChange}
          value={object.nwkKey || ""}
          margin="normal"
          fullWidth
          required
          random
        />
        <AESKeyField
          id="genAppKey"
          label={i18n.t(`${packageNS}:tr000389`)}
          helperText="For LoRaWAN 1.0 devices. This key must only be set when the device implements the remote multicast setup specification / firmware updates over the air (FUOTA). Else leave this field blank."
          onChange={this.onChange}
          value={object.genAppKey || ""}
          margin="normal"
          fullWidth
          random
        />
      </Form>
    );
  }
}


class DeviceKeys extends Component {
  constructor() {
    super();
    this.state = {
      update: false,
    };
    this.onSubmit = this.onSubmit.bind(this);
  }

  componentDidMount() {
    DeviceStore.getKeys(this.props.match.params.devEUI, resp => {
      if (resp === null) {
        this.setState({
          deviceKeys: {
            deviceKeys: {},
          },
        });
      } else {
        this.setState({
          update: true,
          deviceKeys: resp,
        });
      }
    });
  }

  onSubmit(deviceKeys) {
    if (this.state.update) {
      DeviceStore.updateKeys(deviceKeys, resp => {
        this.props.history.push(`/organizations/${this.props.match.params.organizationID}/applications/${this.props.match.params.applicationID}`);
      });
    } else {
      let keys = deviceKeys;
      keys.devEUI = this.props.match.params.devEUI;
      DeviceStore.createKeys(keys, resp => {
        this.props.history.push(`/organizations/${this.props.match.params.organizationID}/applications/${this.props.match.params.applicationID}`);
      });
    }
  }

  render() {
    if (this.state.deviceKeys === undefined) {
      return null;
    }

    return(
      <Grid container spacing={4}>
        <Grid item xs={12}>
          <Card>
            <CardContent>
              {this.props.deviceProfile.macVersion.startsWith("1.0") && <LW10DeviceKeysForm
                submitLabel={i18n.t(`${packageNS}:tr000390`)}
                onSubmit={this.onSubmit}
                object={this.state.deviceKeys.deviceKeys}
              />}
              {this.props.deviceProfile.macVersion.startsWith("1.1") && <LW11DeviceKeysForm
                submitLabel={i18n.t(`${packageNS}:tr000390`)}
                onSubmit={this.onSubmit}
                object={this.state.deviceKeys.deviceKeys}
              />}
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default withRouter(DeviceKeys);
