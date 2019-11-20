import React, { Component } from "react";

import { withStyles } from '@material-ui/core/styles';
import FormControlOrig from "@material-ui/core/FormControl";
import FormLabel from "@material-ui/core/FormLabel";
import TextField from '@material-ui/core/TextField';
import InputLabel from '@material-ui/core/InputLabel';
import Select from '@material-ui/core/Select';
import MenuItem from '@material-ui/core/MenuItem';
import Button from '@material-ui/core/Button';
import Grid from '@material-ui/core/Grid';

import i18n, { packageNS } from '../../i18n';
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import FormControl from "../../components/FormControl";
import AutocompleteSelect from "../../components/AutocompleteSelect";
import NetworkServerStore from "../../stores/NetworkServerStore";

import theme from "../../theme";


const styles = {
  a: {
    color: theme.palette.primary.main,
  },
  formLabel: {
    fontSize: 12,
  },
};


class ExtraChannel extends Component {
  constructor() {
    super();

    this.onChange = this.onChange.bind(this);
    this.onDelete = this.onDelete.bind(this);
  }

  onChange(e) {
    let channel = this.props.channel;
    let field = "";

    if (e.target.id === undefined) {
      field = e.target.name;
    } else {
      field = e.target.id;
    }

    if(field === "spreadingFactorsStr") {
      let sfStr = e.target.value.split(",");
      channel["spreadingFactors"] = sfStr.map((sf, i) => parseInt(sf, 10));
    } 

    if (e.target.type === "number") {
      channel[field] = parseInt(e.target.value, 10);
    } else {
      channel[field] = e.target.value;
    }

    this.props.onChange(channel);
  }

  onDelete(e) {
    e.preventDefault();
    this.props.onDelete();
  }

  render() {
    let spreadingFactorsStr = "";
    if (this.props.channel.spreadingFactorsStr !== undefined) {
      spreadingFactorsStr = this.props.channel.spreadingFactorsStr;
    } else if (this.props.channel.spreadingFactors !== undefined) {
      spreadingFactorsStr = this.props.channel.spreadingFactors.join(", ");
    }

    return(
      <FormControl
        label={
          <span>
            Extra channel {this.props.i + 1} (<a href="#delete" onClick={this.onDelete} className={this.props.classes.a}>delete</a>)
          </span>
        }
      >
        <Grid container spacing={4}>
          <Grid item xs={6}>
            <FormControl>
              <InputLabel htmlFor="modulation" required>{i18n.t(`${packageNS}:tr000118`)}</InputLabel>
              <Select
                value={this.props.channel.modulation || ""}
                onChange={this.onChange}
                inputProps={{
                  name: "modulation",
                }}
              >
                <MenuItem value="LORA">{i18n.t(`${packageNS}:tr000119`)}</MenuItem>
                <MenuItem value="FSK">{i18n.t(`${packageNS}:tr000120`)}</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <FormControl>
              <InputLabel htmlFor="bandwidth" required>{i18n.t(`${packageNS}:tr000121`)}</InputLabel>
              <Select
                value={this.props.channel.bandwidth || ""}
                onChange={this.onChange}
                inputProps={{
                  name: "bandwidth",
                }}
              >
                <MenuItem value={125}>125 {i18n.t(`${packageNS}:tr000122`)}</MenuItem>
                <MenuItem value={250}>250 {i18n.t(`${packageNS}:tr000122`)}</MenuItem>
                <MenuItem value={500}>500 {i18n.t(`${packageNS}:tr000122`)}</MenuItem>
              </Select>
            </FormControl>
          </Grid>
          <Grid item xs={6}>
            <TextField
              id="frequency"
              label="Frequency (Hz)"
              type="number"
              margin="normal"
              value={this.props.channel.frequency || ""}
              onChange={this.onChange}
              required
              fullWidth
            />
          </Grid>
          {this.props.channel.modulation === "LORA" && <Grid item xs={6}>
            <TextField
              id="spreadingFactorsStr"
              label={i18n.t(`${packageNS}:tr000123`)}
              margin="normal"
              value={spreadingFactorsStr}
              onChange={this.onChange}
              placeholder="7, 8, 9, 10, 11, 12"
              helperText={i18n.t(`${packageNS}:tr000124`)}
              inputProps={{
                pattern: "[0-9]+(,[\\s]*[0-9]+)*",
              }}
              required
              fullWidth
            />
          </Grid>}
          {this.props.channel.modulation === "FSK" && <Grid item xs={6}>
            <TextField
              id="bitrate"
              label="Bitrate"
              type="number"
              margin="normal"
              value={this.props.channel.bitrate || ""}
              onChange={this.onChange}
              placeholder="50000"
              required
              fullWidth
            />
          </Grid>}
        </Grid>
      </FormControl>
    );
  }
}

ExtraChannel = withStyles(styles)(ExtraChannel);


class GatewayProfileForm extends FormComponent {
  constructor() {
    super();

    this.addExtraChannel = this.addExtraChannel.bind(this);
    this.getNetworkServerOption = this.getNetworkServerOption.bind(this);
    this.getNetworkServerOptions = this.getNetworkServerOptions.bind(this);
  }

  componentDidMount() {
    super.componentDidMount();

    if (this.props.object !== undefined && this.props.object.channels !== undefined && this.props.object.channelsStr === undefined) {
      let object = this.props.object;
      object.channelsStr = object.channels.join(", ");
      this.setState({
        object: object,
      });
    }
  }

  componentDidUpdate(prevProps) {
    if (prevProps.object !== this.props.object) {
      let object = this.props.object;

      if (object !== undefined && object.channels !== undefined && object.channelsStr === undefined) {
        object.channelsStr = object.channels.join(", ");
      }

      this.setState({
        object: object || {},
      });
    }
  }

  onChange(e) {
    super.onChange(e);

    let object = this.state.object;

    if(e.target.id === "channelsStr") {
      let channelsStr = e.target.value.split(",");
      object["channels"] = channelsStr.map((c, i) => parseInt(c, 10));
    }

    this.setState({
      object: object,
    });
  }

  addExtraChannel() {
    let object = this.state.object;
    if (object.extraChannels === undefined) {
      object.extraChannels = [{modulation: "LORA"}];
    } else {
      object.extraChannels.push({modulation: "LORA"});
    }

    this.setState({
      object: object,
    });
  }

  deleteExtraChannel(i) {
    let object = this.state.object;
    object.extraChannels.splice(i, 1);
    this.setState({
      object: object,
    });
  }

  updateExtraChannel(i, ec) {
    let object = this.state.object;
    object.extraChannels[i] = ec;

    this.setState({
      object: object,
    });
  }

  getNetworkServerOption(id, callbackFunc) {
    NetworkServerStore.get(id, resp => {
      callbackFunc({label: resp.name, value: resp.id});
    });
  }

  getNetworkServerOptions(search, callbackFunc) {
    NetworkServerStore.list(0, 999, 0, resp => {
      const options = resp.result.map((ns, i) => {return {label: ns.name, value: ns.id}});
      callbackFunc(options);
    });
  }

  render() {
    if (this.state.object === undefined) {
      return(<div></div>);
    }

    let extraChannels = [];

    if (this.state.object.extraChannels !== undefined) {
      extraChannels = this.state.object.extraChannels.map((ec, i) => <ExtraChannel key={i} channel={ec} i={i} onDelete={() => this.deleteExtraChannel(i)} onChange={ec => this.updateExtraChannel(i, ec)} />);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.onSubmit}
        extraButtons={<Button onClick={this.addExtraChannel}>{i18n.t(`${packageNS}:tr000116`)}</Button>}
      >
        <TextField
          id="name"
          label={i18n.t(`${packageNS}:tr000042`)}
          margin="normal"
          value={this.state.object.name || ""}
          onChange={this.onChange}
          helperText={i18n.t(`${packageNS}:tr000112`)}
          required
          fullWidth
        />

        <TextField
          id="channelsStr"
          label={i18n.t(`${packageNS}:tr000113`)}
          margin="normal"
          value={this.state.object.channelsStr || ""}
          onChange={this.onChange}
          helperText={i18n.t(`${packageNS}:tr000114`)}
          placeholder="0, 1, 2"
          inputProps={{
            pattern: "[0-9]+(,[\\s]*[0-9]+)*",
          }}
          required
          fullWidth
        />
        {!this.props.update && <FormControlOrig margin="normal" fullWidth>
          <FormLabel className={this.props.classes.formLabel} required>{i18n.t(`${packageNS}:tr000047`)}</FormLabel>
          <AutocompleteSelect
            id="networkServerID"
            label={i18n.t(`${packageNS}:tr000115`)}
            value={this.state.object.networkServerID || ""}
            onChange={this.onChange}
            getOptions={this.getNetworkServerOptions}
          />
        </FormControlOrig>}
        {extraChannels}
      </Form>
    );
  }
}

export default withStyles(styles)(GatewayProfileForm);
