import React, { useEffect, useState, useContext } from 'react';
import { GroupContext } from './App';
import { db } from './api';
import { Card, Elevation, H3, Text } from '@blueprintjs/core';

const Chats = props => {
    const [chats, setChats] = useState([]);
    const group = useContext(GroupContext);

    useEffect(() => {
        db.collection('messages')
            .where('chat', '==', group)
            .orderBy('timestamp', 'desc')
            .limit(100)
            .get()
            .then(qs => setChats(qs.docs.map(doc => doc.data())));
    }, [group]);

    if (!group) {
        return (
            <Card elevation={Elevation.TWO}>
                <H3>Chatit</H3>
                Valitse keskustelu yläpalkista!
            </Card>
        );
    }

    console.log(chats);

    return (
        <Card elevation={Elevation.TWO}>
            <H3>Chatit</H3>
            {!chats && 'Ladataan...'}
            <div style={{ maxHeight: 600, overflow: 'scroll', padding: 16 }}>
                {([...chats].reverse()).map((msg, i) => (
                    <Card key={i} style={{ marginTop: 16 }} elevation={Elevation.ONE}>
                        <strong>{msg.author}</strong>
                        <Text className="bp3-text-small bp3-text-muted">{new Date(msg.timestamp * 1000).toLocaleString()}</Text>
                        {msg.message || <em className="bp3-text-muted">Ei viestiä</em>}
                    </Card>
                ))}
            </div>
        </Card>
    );
};

export default Chats;
