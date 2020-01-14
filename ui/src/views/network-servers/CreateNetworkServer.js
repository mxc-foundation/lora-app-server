import React, { Component } from "react";
import { Link, withRouter } from 'react-router-dom';
import { Breadcrumb, BreadcrumbItem, Card } from 'reactstrap';

import { withStyles } from "@material-ui/core/styles";
import i18n, { packageNS } from '../../i18n';
import NetworkServerStore from "../../stores/NetworkServerStore";
import TitleBar from "../../components/TitleBar";
import NetworkServerForm from "./NetworkServerForm";

import breadcrumbStyles from "../common/BreadcrumbStyles";
import Admin from "../../components/Admin";

const localStyles = {};

const styles = {
  ...breadcrumbStyles,
  ...localStyles
};

class CreateNetworkServer extends Component {
  constructor() {
    super();
    this.onSubmit = this.onSubmit.bind(this);
  }

  onSubmit(networkServer) {
    NetworkServerStore.create(networkServer, resp => {
      this.props.history.push("/network-servers");
    });
  }

  render() {
    const { classes } = this.props;

    return(
      <React.Fragment>
        <TitleBar>
          <Breadcrumb className={classes.breadcrumb}>
            <Admin>
              <BreadcrumbItem className={classes.breadcrumbItem}>Control Panel</BreadcrumbItem>
            </Admin>
            <BreadcrumbItem><Link className={classes.breadcrumbItemLink} to={`/network-servers`}>{i18n.t(`${packageNS}:tr000040`)}</Link></BreadcrumbItem>
            <BreadcrumbItem active>{i18n.t(`${packageNS}:tr000277`)}</BreadcrumbItem>
          </Breadcrumb>
        </TitleBar>

        <Card className="card-box shadow-sm" style={{ minWidth: "25rem" }}>
          <NetworkServerForm
            onSubmit={this.onSubmit}
            submitLabel={i18n.t(`${packageNS}:tr000041`)}
          />
        </Card>
      </React.Fragment>
    );
  }
}

export default withStyles(styles)(withRouter(CreateNetworkServer));
