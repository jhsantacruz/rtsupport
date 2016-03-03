import React, {Component} from 'react';

class User extends Component{
  render(){
    return(
      <li>
        {this.props.user.name}
      </li>
    )
  }
}

User.propTypes = {
  User: React.PropTypes.object.isRequired
}

export default User
