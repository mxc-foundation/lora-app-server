import React, { Component } from "react";
import { withRouter } from 'react-router-dom';
import { Card } from 'reactstrap';
import Loader from "../../components/Loader";
import i18n, { packageNS } from '../../i18n';
import NetworkServerStore from "../../stores/NetworkServerStore";
import NetworkServerForm from "./NetworkServerForm";



class UpdateNetworkServer extends Component {
  constructor() {
    super();
    this.state = {
      loading: false,
    }
    this.onSubmit = this.onSubmit.bind(this);
  }

 onSubmit = async (networkServer) => {
    this.setState({ loading: true });
    let res = await NetworkServerStore.update(networkServer);
  
    this.props.history.push("/network-servers");
    this.setState({ loading: false })
  }

  render() {
    return(
      <React.Fragment>
        <Card className="card-box shadow-sm">
          {this.state.loading && <Loader />}
          
          <NetworkServerForm
            object={this.props.networkServer}
            version={this.props.version}
            region={this.props.region}
            onSubmit={this.onSubmit}
            submitLabel={i18n.t(`${packageNS}:tr000066`)}
          />
        </Card>
      </React.Fragment>
    );
  }
}

export default withRouter(UpdateNetworkServer);
