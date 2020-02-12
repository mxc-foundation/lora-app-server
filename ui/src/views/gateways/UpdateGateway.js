import React, { Component } from "react";
import { withRouter } from "react-router-dom";

import i18n, { packageNS } from "../../i18n";
import GatewayStore from "../../stores/GatewayStore";
import Loader from "../../components/Loader";
import GatewayForm from "./GatewayForm";

class UpdateGateway extends Component {
  constructor() {
    super();
    this.onSubmit = this.onSubmit = this.onSubmit.bind(this);
    this.state = {
      loading: false
    };
  }

  onSubmit(gateway, config, classBConfig) {
    this.setState({ loading: true });
    GatewayStore.update(gateway, resp => {
      this.setState({ loading: false });
      this.props.history.push(
        `/organizations/${this.props.match.params.organizationID}/gateways`
      );
    });

    // TODO - save config using setConfig api call
    console.log(config);

    // TODO - save class B configuration
    console.log(classBConfig);
  }

  render() {
    return (
      <div className="position-relative">
        {this.state.loading && <Loader />}
        <GatewayForm
          submitLabel={i18n.t(`${packageNS}:tr000614`)}
          object={this.props.gateway}
          onSubmit={this.onSubmit}
          update={true}
          match={this.props.match}
        ></GatewayForm>
      </div>
    );
  }
}

export default withRouter(UpdateGateway);
