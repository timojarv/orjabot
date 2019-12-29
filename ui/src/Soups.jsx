import React, { useEffect, useState, useContext } from 'react';
import { GroupContext } from './App';
import { db } from './api';
import { Card, Elevation, H3, Button, Icon, Classes } from '@blueprintjs/core';

const ymd = ts => ts.toISOString().split('T')[0];
const parseymd = str => Date.parse(str).valueOf() / 1000 + 24 * 3600 - 1;

const renderStars = (amount, onClick) => {
    return (
        <div style={{ margin: '8px 0' }}>
            {Array(5).fill(true).map((_, i) =>
                <Icon
                    style={{ marginRight: 4, cursor: onClick ? 'pointer' : 'initial' }}
                    intent="primary"
                    key={i}
                    icon={amount > i ? 'star' : 'star-empty'}
                    onClick={() => onClick(i + 1)}
                />
            )}
        </div>
    );
};

const SoupView = props => {
    const { soup, onEdit } = props;
    return (
        <Card style={{ marginTop: 16 }} elevation={Elevation.ONE}>
            <Button onClick={() => onEdit(soup.id)} style={{ float: 'right' }} icon="edit" />
            <strong>{soup.name}</strong> <br />
            {renderStars(soup.index)}
            <p className="bp3-text-muted">{new Date(soup.date * 1000).toUTCString()}</p>
        </Card>
    );
};

const SoupEditView = props => {
    const { onSave, onCancel } = props;
    const [soup, setSoup] = useState(props.soup);
    return (
        <Card style={{ marginTop: 16 }} elevation={Elevation.TWO}>
            <Button onClick={() => onSave(soup)} style={{ float: 'right', marginLeft: 8 }} icon="tick" />
            <Button onClick={onCancel} style={{ float: 'right' }} icon="cross" />
            <input className={Classes.INPUT} style={{ width: 200 }} value={soup.name} onChange={e => setSoup({ ...soup, name: e.target.value })} />
            {renderStars(soup.index, n => setSoup({ ...soup, index: n }))}
            <input className={Classes.INPUT} type="date" value={ymd(new Date(soup.date * 1000))} onChange={e => setSoup({ ...soup, date: parseymd(e.target.value) })} />
        </Card>
    );
};

const fetchData = (group, setter) => db.collection('keitot')
    .where('group', '==', group)
    .where('date', '>=', Date.now() / 1000)
    .get()
    .then(qs => setter(qs.docs.map(doc => ({ id: doc.id, ...doc.data() }))));

const Soups = props => {
    const [soups, setSoups] = useState([]);
    const [editing, setEditing] = useState(false);
    const group = useContext(GroupContext);

    useEffect(() => {
        fetchData(group, setSoups);
    }, [group]);

    const handleSave = soup => {
        const { id, ...data } = soup;
        db.collection('keitot').doc(id).set(data)
            .then(() => setEditing(false))
            .then(() => fetchData(group, setSoups));
    };

    if (!group) {
        return (
            <Card elevation={Elevation.TWO}>
                <H3>Keitot</H3>
                Valitse keskustelu yläpalkista!
            </Card>
        );
    }

    console.log(soups)

    return (
        <Card elevation={Elevation.TWO}>
            <H3>Keitot</H3>
            {!soups && 'Ladataan...'}
            {soups.map((soup, i) => (
                editing === soup.id
                    ? <SoupEditView onSave={handleSave} onCancel={() => setEditing(false)} key={i} soup={soup} />
                    : <SoupView onEdit={id => setEditing(id)} key={i} soup={soup} />
            ))}
        </Card>
    );
};

export default Soups;
