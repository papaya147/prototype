import React, { Component } from 'react';
import { View, Text } from 'react-native';
import socketIOClient from 'socket.io-client';
import { StyleSheet } from 'react-native';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      message: '',
      data: 0,
      time: Date.now()
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

    setInterval(() => {
      this.setState({ data: Math.floor(Math.random() * 100), time: Date.now() })
    }, 3000)

    // You can also emit messages to the server
    // socket.emit('sendMessage', { message: 'Hello, server!' });
  }

  render() {
    return (
      <View style={styles.container}>
        <Text>Sending Message: data = {this.state.data}, time = {this.state.time}</Text>
        <Text>WebSocket Message: {this.state.message}</Text>
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

export default App;
