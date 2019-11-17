import React, { Component } from "react";
import Select from "react-select";
import { withStyles } from "@material-ui/core/styles";
import classNames from "classnames";
import { DEFAULT_LANGUAGE, SUPPORTED_LANGUAGES } from "../i18n";
import SessionStore from "../stores/SessionStore";
import FlagIcon from "./FlagIcon";

const styles = {
  languageWrapper: {
    display: "inline-flex"
  },
  languageIcon: {
    display: "inline-block",
  },
  languageSelection: {
    display: "inline-block"
  }
};

const customStyles = {
  control: (base, state) => ({
    ...base,
    color: "#FFFFFF",
    width: "180px",
    margin: 20,
    // match with the menu
    borderRadius: state.isFocused ? "3px 3px 0 0" : 3,
    // Overwrittes the different states of border
    borderColor: state.isFocused ? "#00FFD9" : "white",
    // Removes weird border around container
    boxShadow: state.isFocused ? null : null,
    "&:hover": {
      // Overwrittes the different states of border
      borderColor: state.isFocused ? "#00FFD9" : "white"
    }
  }),
  menu: base => ({
    ...base,
    background:"#101c4a",
    // override border radius to match the box
    borderRadius: 0,
    // kill the gap
    marginTop: 0,
    // paddingLeft: 20,
    // paddingRight: 20,
  }),
  menuList: base => ({
    ...base,
    background: "#1a2d6e",
    // kill the white space on first and last option
    paddingTop: 0,
  }),
  option: base => ({
    ...base,
    // kill the white space on first and last option
    padding: "10px",
    maxWidth: 229,
    whiteSpace: "nowrap", 
    overflow: "hidden",
    textOverflow: "ellipsis"
  }),
};

class WithPromises extends Component {
  constructor() {
    super();

    this.state = {
      selectedOption: null,
      options: []
    };
  } 

  componentDidMount() {
    let selectedOption = null;

    const language = SessionStore.getLanguage();

    if (!language || !language.label) {
      selectedOption = DEFAULT_LANGUAGE;
    } else if (language.label && language.value && language.code) {
      selectedOption = {
        label: language.label,
        value: language.value,
        code: language.code
      };
    }

    this.setState({
      selectedOption,
      options: SUPPORTED_LANGUAGES
    });
  }

  onChangeLanguage = selectedOption => {
    if (
      selectedOption !== null && selectedOption.label !== null &&
      selectedOption.value !== null && selectedOption.code !== null
    ) {
      this.setState({
        selectedOption
      });
  
      this.props.onChangeLanguage({
        label: selectedOption.label,
        value: selectedOption.value,
        code: selectedOption.code
      });
    }
  }

  render() {
    const { selectedOption } = this.state;

    return (
      <div className={classNames(this.props.classes.languageWrapper)}>
        {
          selectedOption && selectedOption.code
          ? (
            <FlagIcon
              className={classNames(this.props.classes.languageIcon)}
              code={selectedOption.code}
              size="2x"
            />
          ) : null
        }
        <Select
          className={classNames(this.props.classes.languageSelection)}
          styles={customStyles}
          theme={(theme) => ({
            ...theme,
            borderRadius: 4,
            colors: {
              primary25: "#00FFD950",
              primary: "#00FFD950",
            },
          })}
          placeholder="Select Language"
          onChange={this.onChangeLanguage}
          options={SUPPORTED_LANGUAGES}
          value={selectedOption}
        />
      </div>
    );
  }
}

export default withStyles(styles)(WithPromises);
