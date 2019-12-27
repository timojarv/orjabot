import React, { useEffect, useState, useContext } from 'react';
import { Menu, Popover, Position, Button } from '@blueprintjs/core';
import { db } from './api';
import { GroupContext } from './App';

const GroupSelector = props => {

    const [groups, setGroups] = useState([]);

    const selected = useContext(GroupContext);

    useEffect(() => {
        db.collection('groups').get()
            .then(qs => setGroups(qs.docs.map(doc => doc.id)));
    }, [setGroups])

    return (
        <Popover content={
            <Menu>
                {groups.map(group => (
                    <Menu.Item onClick={() => props.onChange(group)} key={group} icon={group > 0 ? 'person' : 'people'} text={group} />
                ))}
            </Menu>}
            position={Position.BOTTOM_RIGHT}>
            <Button minimal icon="caret-down" text={selected || 'Keskustelu'} />
        </Popover>
    );
};

export default GroupSelector;
