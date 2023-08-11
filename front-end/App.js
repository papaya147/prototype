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
    const socket = socketIOClient('ws://65.1.194.238:4000');

    socket.on('connect', () => {
      console.log('Connected to server');
    });

    socket.on('ack', (data) => {
      this.setState({ message: JSON.parse(data).data, time: Date.now() });
    });

    setInterval(() => {
      this.setState({ data: Math.floor(Math.random() * 100), time: Date.now() })
      socket.emit('telemetry', JSON.stringify({ data: this.state.data, time: this.state.time }))
    }, 3000)

    // You can also emit messages to the server
    // socket.emit('sendMessage', { message: 'Hello, server!' });
  }

  render() {
    return (
      <View style={styles.container}>
        <Text>Sending Message: data = {this.state.data}, time = {this.state.time}</Text>
        <Text>WebSocket Message: data = {this.state.message}, time = {this.state.time}</Text>
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
