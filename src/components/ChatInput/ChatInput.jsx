import React,{Component} from 'react';
import './ChatInput.scss'

export const ChatInput = (props) => {

    return (
        <div className='ChatInput'>
        <input onKeyDown={this.props.send} placeholder="Type a message... Hit Enter to Send"/>
      </div>
    )
}