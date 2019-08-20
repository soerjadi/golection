import React, { Component } from 'react';
import './IndexPage.scss';

export class IndexPage extends Component {
    constructor(props) {
        super(props);
        this.state = {
            identity: ""
        };

        this.handleSubmit = this.handleSubmit.bind(this)
        this.handleChange = this.handleChange.bind(this)
    }
    handleChange(event) {
        this.setState({identity: event.target.value});
    }
    handleSubmit(event) {
        alert(this.state.identity);
        event.preventDefault();
    }
    render() {
        return (
            <div className="container">
                <div className="row align-items-center">
                    <div className="col"></div>
                    <div className="col">
                        <form onSubmit={this.handleSubmit}>
                            <div className="form-group">
                                <label>NIK</label>
                                <input type="text" name="identity-id" className="form-control" value={this.state.identity} onChange={this.handleChange} placeholder="xxxxxxxxxxxxx" />
                            </div>
                            <div className="form-group">
                                <input type="submit" name="submit" value="Masuk" className="btn btn-info" />
                            </div>
                        </form>
                    </div>
                    <div className="col"></div>
                </div>
                
            </div>
        )
    }
}

export default (IndexPage);