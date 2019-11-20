import React, { Component } from "react";
import { withRouter } from 'react-router-dom';

import Grid from '@material-ui/core/Grid';
import Card from '@material-ui/core/Card';
import CardContent from "@material-ui/core/CardContent";
import TextField from '@material-ui/core/TextField';

import TitleBar from "../../components/TitleBar";
import TitleBarTitle from "../../components/TitleBarTitle";
import UserStore from "../../stores/UserStore";
import FormComponent from "../../classes/FormComponent";
import Form from "../../components/Form";
import i18n, { packageNS } from '../../i18n';


class PasswordForm  extends FormComponent {
  render() {
    if (this.state.object === undefined) {
      return(<div></div>);
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        onSubmit={this.onSubmit}
      >
        <TextField
          id="password"
          label={i18n.t(`${packageNS}:tr000004`)}
          type="password"
          margin="normal"
          value={this.state.object.password || ""}
          onChange={this.onChange}
          required
          fullWidth
        />
      </Form>
    );
  }
}


class ChangeUserPassword extends Component {
  constructor() {
    super();
    this.state = {};

    this.onSubmit = this.onSubmit.bind(this);
  }

  componentDidMount() {
    UserStore.get(this.props.match.params.userID, resp => {
      this.setState({
        user: resp,
      });
    });
  }

  onSubmit(password) {
    UserStore.updatePassword(this.props.match.params.userID, password.password, resp => {
      this.props.history.push("/");
    });
  }

  render() {
    if (this.state.user === undefined) {
      return(<div></div>);
    }

    return(
      <Grid container spacing={4}>
        <TitleBar>
          <TitleBarTitle title={i18n.t(`${packageNS}:tr000036`)} />
          <TitleBarTitle title="/" />
          <TitleBarTitle title={this.state.user.user.username} />
          <TitleBarTitle title="/" />
          <TitleBarTitle title={i18n.t(`${packageNS}:tr000038`)} />
        </TitleBar>

        <Grid item xs={12}>
          <Card>
            <CardContent>
              <PasswordForm
                submitLabel={i18n.t(`${packageNS}:tr000022`)}
                onSubmit={this.onSubmit}
              />
            </CardContent>
          </Card>
        </Grid>
      </Grid>
    );
  }
}

export default withRouter(ChangeUserPassword);
