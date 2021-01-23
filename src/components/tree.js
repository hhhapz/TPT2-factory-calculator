import React, { useEffect } from 'react'
import CheckboxTree from 'react-checkbox-tree'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import 'react-checkbox-tree/lib/react-checkbox-tree.css';
import { faChevronDown, faChevronRight, faPlusSquare } from '@fortawesome/free-solid-svg-icons'

function Tree({ item, checked, setChecked }) {
    const [expanded, setExpanded] = React.useState([])
    const [nodes, setNodes] = React.useState([])

    useEffect(() => {
        if (!item) setNodes([])
        else {
            const [tree] = item.tree()
            setNodes([tree])
            setExpanded([item.id])
        }
    }, [item, setNodes])

    return <div id="recipe">
        <h2>Recipe:</h2>

        <CheckboxTree nodes={nodes} showNodeIcon={false} nativeCheckboxes={true}
            checked={checked} expanded={expanded}
            onCheck={checked => { setChecked(checked) }}
            onExpand={expanded => setExpanded(expanded)}
            icons={{
                expandClose: <FontAwesomeIcon icon={faChevronRight} />,
                expandOpen: <FontAwesomeIcon icon={faChevronDown} />,
                expandAll: <FontAwesomeIcon icon={faPlusSquare} />,
            }}
        ></CheckboxTree>
    </div>

}

export default Tree