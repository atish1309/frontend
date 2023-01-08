import React, { Component } from 'react';
import './ChatHistory.scss';
import Message from '../Message/Message';

export default ChatHistory =()=>{
    return(
        <div className='ChatHistory'>
        <h2>Chat History</h2>
        {messages}
      </div>
    )
}