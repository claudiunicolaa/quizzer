import React, {Component} from 'react';
import PropTypes from 'prop-types';
import {AppSwitch} from "@coreui/react";
import {ListGroup, ListGroupItem} from "reactstrap";
import {connect} from "react-redux";
import {setQuestionScore} from "../../../redux/actions";

class CheckboxQuestion extends Component {
  static propTypes = {
    question: PropTypes.shape({
      ID: PropTypes.number.isRequired,
      Text: PropTypes.string.isRequired,
      ChoiceAnswers: PropTypes.arrayOf(PropTypes.shape({
        Text: PropTypes.string.isRequired
      })).isRequired,
      Score: PropTypes.number.isRequired,
    }).isRequired
  };

  state = {
    Score: 0
  };

  componentWillReceiveProps(nextProps, nextContext) {
    if (nextProps.question !== undefined) {
      this.setState({
        Score: this.props.question.Score
      })
    }
  }

  setScore = (e) => {
    this.props.setQuestionScore(
      this.props.question,
      e.target.checked ? 100 : 0
    );

    this.setState({
      Score: e.target.checked ? 100 : 0
    });
  };

  render() {
    return (
      <div>
        <div>
          <div className="float-right">
            <AppSwitch
              className={'mx-1'}
              color={'success'}
              outline={'alt'}
              defaultChecked={this.props.question.Score === 100}
              onChange={this.setScore}
              label
              dataOn={'\u2713'}
              dataOff={'\u2715'} />
          </div>
          <div dangerouslySetInnerHTML={{__html: this.props.question.Text}} />
          <small>
            <span style={{color: "#4DBD74"}}>GREEN</span>
            {' '}answers are the correct ones
          </small>
          <br/>
          <small><i className="fa fa-check-square-o" /> answers are the ones the user selected</small>
          <br/>
          <small><i className="fa fa-square-o" /> answers are the ones the user did not select</small>
        </div>
        <div>
          <ListGroup>
            {this.props.question.ChoiceAnswers.map((a, i) => (
              <ListGroupItem
                style={{color: a.IsCorrect ? "#4DBD74" : "inherit"}}
                key={i}>
                {a.IsSelected
                  ? <i className="fa fa-check-square-o" />
                  : <i className="fa fa-square-o" />
                }
                {' ' + a.Text}
              </ListGroupItem>
            ))}
          </ListGroup>
        </div>
      </div>
    );
  }
}

export default connect(
  null,
  {setQuestionScore}
)(CheckboxQuestion);
