import React, { Component } from "react";
import { Link } from "react-router-dom";

import classNames from "classnames";

import Typography from '@material-ui/core/Typography';
import { withStyles } from '@material-ui/core/styles';

import theme from "../theme";

const styles = {
  title: {
    marginTop: theme.spacing.unit,
    marginBottom: theme.spacing.unit,
    marginRight: theme.spacing.unit,
    float: "left",
  },

  link: {
    textDecoration: "none",
    color: theme.palette.primary.main,
  },
  padding: {
    padding: 30,
  },
};


class TypographyStyled extends Component {
  render() {
    let component = null;

    if (this.props.to !== undefined) {
      component = Link;
    }
    
    const styleArr = [];
    if(this.props.styles.length > 0){
        for(let i=0; i < this.props.styles.length; i++){
            styleArr.push(this.props.classes[this.props.styles[i]]);
        }
    }

    let combinedStyles = classNames(...styleArr);
    
    return(
      <Typography variant={this.props.variant} 
                className={combinedStyles} 
                to={this.props.to} 
                component={component}
            >
        {this.props.text}
      </Typography>
    );
  }
}

export default withStyles(styles)(TypographyStyled);
