import React, { Component } from "react";
import { withRouter, Link } from "react-router-dom";

import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';
import Button from '@material-ui/core/Button';
import TitleBarTitle from '../../components/TitleBarTitle';

import AppBar from '@material-ui/core/AppBar';
import Toolbar from '@material-ui/core/Toolbar';
import IconButton from '@material-ui/core/IconButton';
//import MenuIcon from 'mdi-material-ui/Server';
import Password from '../../components/TextfileForPassword'
import { 
  Map,
  Marker,
  Popup,
  LayersControl,
  LayerGroup
 } from 'react-leaflet';

import Form from "../../components/Form";
import FormComponent from "../../classes/FormComponent";
import SessionStore from "../../stores/SessionStore";
import theme from "../../theme";
import MapTileLayerCluster from "../../components/MapTileLayerCluster";
//import { isAbsolute } from "path";
//import { NONAME } from "dns";
//import { relative } from "path";
//const DURATION = 550;
//const COLOR = 'rgba(121,244,218,0.5)';

const styles = {
  textField: {
    width: "100%",
  },
  link: {
    "& a": {
      color: theme.palette.textSecondary.main,
      textDecoration: "none",
    },
  },
  padding: {
    paddingTop: 115,
  },
  z1000: {
    zIndex: 1000
  },
  loginFormStyle: {
    backgroundColor: '#09006E',
    padding: '24px',
    position: 'absolute',
    width: 400,
    top: 90,
    right: 40,
    display: 'flex',
    justifyContent: 'center',
    alignItems: 'center',
    height: 400,
  },
  logo: {
    height: 90,
    marginLeft: 'auto',
    opacity: '0.7',
  },
  logoSection: {
    display: 'flex'
  },
  root: {
    flexGrow: 1,
    position: 'absolute',
    top: 0,
    height: 84, 
    left: 0,
    right: 0,
    zIndex: 1
  },
  menuButton: {
    marginRight: theme.spacing(2),
  },
  title: {
    flexGrow: 1,
  },
  appBar: {
    backgroundColor: theme.palette.secondary.main,
  },
  logo:{
    height: 50 
  }
};

class LoginForm extends FormComponent {
  constructor(props) {
    super(props);
    this.handleChange = this.handleChange.bind(this);
  }

  handleChange = (event) => {
    this.state.object.password = event;
  };

  render() {
    if (this.state.object === undefined) {
      return null;
    }
    
    const extraButtons = [
      <Button 
        variant="outlined"
        color="inherit"
        component={Link} 
        to={`/registration`} 
        type="button" 
        disabled={false}>Register</Button>
    ]
    let demoUsername = "";
    let demoPassword = "";
    let helpText = "";
    if(window.location.origin.includes(process.env.REACT_APP_DEMO_HOST_SERVER)){
      demoUsername = process.env.REACT_APP_DEMO_USER;
      demoPassword = process.env.REACT_APP_DEMO_USER_PASSWORD;
      helpText = "You can access right now as a demo user.";
    }

    return(
      <Form
        submitLabel={this.props.submitLabel}
        extraButtons={extraButtons}
        onSubmit={this.onSubmit}
      >
        <TextField
          id="username"
          label="E-Mail"
          margin="normal"
          value={this.state.object.username === undefined 
            ? this.state.object.username = demoUsername 
            : this.state.object.username }
          autoComplete='off'
          onChange={this.onChange}
          fullWidth
        />
        <Password handleChange={this.handleChange} demoPassword={demoPassword} helpText={helpText}/>
        <TitleBarTitle component={Link} to={`#`} title="FORGOT MY PASSWORD" />
      </Form>
    );
  }
}


class Login extends Component {
  constructor() {
    super();

    this.state = {
      registration: null,
      open: true,
      accessOn: false
    };

    this.onSubmit = this.onSubmit.bind(this);
  }

  componentDidMount() {
    SessionStore.logout(() => {});

    SessionStore.getBranding(resp => {
      if (resp.registration !== "") {
        this.setState({
          registration: resp.registration,
        });
      }
    });
  }

  onSubmit(login) {
    SessionStore.login(login, () => {
      this.props.history.push("/");
    });
  }

  onClick = () => {
    this.setState(function(prevState) {
			return {accessOn: !prevState.accessOn};
		});
  }

  render() {
    const style = {
      position: 'absolute',
      top: 84,
      bottom: 0,
      left: 0,
      right: 0,
      zIndex: 1
    };

    let position = [];
    
    position = [51,13];
    
    return(
      <>
      <div className={this.props.classes.root}>
        <AppBar position="static" className={this.props.classes.appBar}>
          <Toolbar>
            <div className={this.props.logoSection}>
              <img src="/logo/logo_mx.png" className={this.props.classes.logo} alt="LoRa Server" />
            </div>
            <IconButton edge="start" className={this.props.classes.menuButton} color="inherit" aria-label="menu">
              {/* <MenuIcon /> */}
            </IconButton>
            <Typography variant="h6" className={this.props.classes.title}></Typography>
            <Button variant="outlined"
                    color="inherit"
                    onClick={this.onClick}
            >ACCESS</Button>
          </Toolbar>
        </AppBar>
      </div>
        <Map center={position} zoom={6} style={style} animate={true} scrollWheelZoom={false}>
          <MapTileLayerCluster />
        </Map>
        {this.state.accessOn && <div className={this.props.classes.padding + ' ' + this.props.classes.z1000}>
          <div className={this.props.classes.loginFormStyle}>
            <LoginForm
              submitLabel="Login"
              onSubmit={this.onSubmit}
              logo={this.props.classes.logo}
            />
          </div>
          {this.state.registration && <div>
            <Typography className={this.props.classes.link} dangerouslySetInnerHTML={{__html: this.state.registration}}></Typography>
          </div>}
        </div>}
      </>
    );
  }
}

export default withStyles(styles)(withRouter(Login));
