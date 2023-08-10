import React, { Component } from 'react';
import { View, Text } from 'react-native';
import socketIOClient from 'socket.io-client';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: '',
    };
  }

  componentDidMount() {
    // Replace 'http://your-socket-server-url' with your actual WebSocket server URL
    const socket = socketIOClient('http://your-socket-server-url');

    socket.on('connect', () => {
      console.log('Connected to server');
    });

    socket.on('message', (data) => {
      this.setState({ message: data.message });
    });

    // You can also emit messages to the server
    // socket.emit('sendMessage', { message: 'Hello, server!' });
  }

  render() {
    return (
      <View>
        <Text>WebSocket Message: {this.state.message}</Text>
      </View>
    );
  }
}

export default App;
