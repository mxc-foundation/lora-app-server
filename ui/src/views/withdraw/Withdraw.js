import React, { Component } from "react";
import { withRouter, Link } from "react-router-dom";

import { Row, Col, Card, Button, Breadcrumb, BreadcrumbItem, FormGroup, Label, Input } from 'reactstrap';
import { withStyles } from "@material-ui/core/styles";
import localStyles from "./WithdrawStyle"
import i18n, { packageNS } from "../../i18n";
import NumberFormat from 'react-number-format';
import WithdrawHistory from "./WithdrawHistory";
import WithdrawForm from "./WithdrawForm";


import breadcrumbStyles from "../common/BreadcrumbStyles";
import Modal from './Modal';
import { MAX_DATA_LIMIT } from '../../util/pagination';
import TitleBar from "../../components/TitleBar";
import AdvancedTable from "../../components/AdvancedTable";
import Loader from "../../components/Loader";
import WithdrawStore from "../../stores/WithdrawStore";
import MoneyStore from "../../stores/MoneyStore";
import WalletStore from "../../stores/WalletStore";

const styles = {
  ...breadcrumbStyles,
  ...localStyles
};

class Withdraw extends Component {
  constructor(props) {
    super(props);
    this.state = {
      data: [],
      stats: {},
      totalSize: 0,
      nsDialog: false
    }
  }
  
  loadData = () => {

    const orgId = this.props.match.params.organizationID;

    MoneyStore.getActiveMoneyAccount(0, orgId, resp => {
      console.log('resp', resp);
      this.setState({
        activeAccount: resp.activeAccount,
      });
    });
  }

  componentDidMount() {
    //this.loadData();
  }

  componentDidUpdate(prevProps, prevState) {
    if (prevState !== this.state && prevState.data !== this.state.data) {

    }
  }

  render() {
    const { classes } = this.props;
    const currentOrgID = this.props.organizationID || this.props.match.params.organizationID;

    return (

      <React.Fragment>
        {this.state.nsDialog && <Modal
          title={i18n.t(`${packageNS}:menu.withdraw.confirm_modal_title`)}
          context={(this.state.status) ? i18n.t(`${packageNS}:menu.withdraw.confirm_text`) : i18n.t(`${packageNS}:menu.withdraw.deny_text`)}
          status={this.state.status}
          row={this.state.row}
          handleChange={this.handleChange}
          closeModal={() => this.setState({ nsDialog: false })}
          callback={() => { this.confirm(this.state.row, this.state.status) }} />}
        <TitleBar>

          <Breadcrumb className={classes.breadcrumb}>
            <BreadcrumbItem>
              <Link
                className={classes.breadcrumbItemLink}
                to={`/organizations`}
                onClick={() => {
                  // Change the sidebar content
                  this.props.switchToSidebarId('DEFAULT');
                }}
              >
                {i18n.t(`${packageNS}:tr000049`)}
              </Link>
            </BreadcrumbItem>
            <BreadcrumbItem className={classes.breadcrumbItem}>{i18n.t(`${packageNS}:tr000084`)}</BreadcrumbItem>
            <BreadcrumbItem active>{i18n.t(`${packageNS}:menu.withdraw.withdraw`)}</BreadcrumbItem>
          </Breadcrumb>
        </TitleBar>
        <WithdrawForm />
        <Row>
          <Col>
            <Card className="card-box shadow-sm">
              {/* {this.state.loading && <Loader />} */}
              <WithdrawHistory />
            </Card>
          </Col>
        </Row>
      </React.Fragment>
    );
  }
}

export default withStyles(styles)(withRouter(Withdraw));
