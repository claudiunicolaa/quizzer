import {Card, CardHeader} from "reactstrap";
import React, {Component} from "react";
import CreateStep1 from "./CreateStep1";
import CreateStep2 from "./CreateStep2";
import {connect} from "react-redux";
import {createQuizTemplate} from "../../../redux/actions";

class Create extends Component {
  defaultState = {
    step: 1,
    quiz: {
      Name: '',
      Questions: []
    }
  };

  state = this.defaultState;

  advanceToStep2 = (text) => {
    this.setState({
      step: 2,
      quiz: {
        Name: text,
        Questions: []
      }
    })
  };

  stop = (qIds) => {
    const quiz = this.state.quiz;
    quiz.Questions = qIds.map(id => ({"ID": id}));

    this.props.createQuizTemplate(quiz)
      .then(() => this.setState(this.defaultState))
  };

  render() {
    return (
      <Card>
        <CardHeader>
          <i className="fa fa-plus-circle text-success" />
          <strong>Create quiz</strong>
          <small> Wizzard</small>
        </CardHeader>
        {this.state.step === 1 &&
        <CreateStep1
          advance={this.advanceToStep2}/>
        }
        {this.state.step === 2 &&
        <CreateStep2
          back={() => this.setState({step: 1})}
          advance={this.stop}/>
        }
      </Card>
    );
  }
}

export default connect(
  null,
  {createQuizTemplate}
)(Create);
