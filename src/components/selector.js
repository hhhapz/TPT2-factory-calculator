import React from "react"
import { config } from "../recipes"
import Item from "../item"

function Selector({ setItem }) {
    const [tier, setTier] = React.useState("")
    const [machine, setMachine] = React.useState("")

    const updateOption = setter => event => {
        setter(event.target.value)
    }

    const calculate = () => {
        if (tier === "" || machine === "") return

        const item = new Item(machine, tier)
        setItem(item)
    }

    return <div id="selector">
        <label htmlFor="tier">Tier</label>
        <select id="tier" onChange={updateOption(setTier)}>
            <option disabled selected value="">Select Tier</option>
            {[...Array(config.tiers).keys()].map(i => <option key={i} value={i + 1}>T{i + 1}</option>)}
        </select>

        <label htmlFor="machine">Machine</label>
        <select id="machine" onChange={updateOption(setMachine)}>
            <option disabled selected value="">Select Machine</option>
            {config.machines.map((machine, i) => <option key={i} value={machine}>{machine}</option>)}
        </select>

        <button onClick={calculate}>Calculate</button>
    </div>
}

export default Selector