import React from 'react'
import './App.css'
import Selector from './components/selector'
import Table from './components/table'
import Tree from './components/tree'

function App() {
    const [item, setItem] = React.useState(null)
    const [checked, setChecked] = React.useState([])

    return (
        <div className="App">
            <h1>Perfect Tower 2 - Cost Calculator</h1>

            <div>
                <Selector setItem={setItem} />

                <Tree item={item} checked={checked} setChecked={setChecked}></Tree>
            </div>

            <Table item={item} checked={checked}></Table>
        </div>
    )
}

export default App
