import { React } from 'react'

function Table({ item, checked }) {

    const totalDust = item && item.totalDust(checked)
    return <div id="summary">
        <h2>Summary:</h2>

        {item &&
            <div id="tables">
                <div>
                    <h3>Item breakdown:</h3>
                    <table>
                        <tbody>
                            <tr>
                                <th>Item</th>
                                <th>Tier</th>
                                <th>Amount</th>
                            </tr>
                            {item.summary(true, checked).map(([name, tier, count], i) =>
                                <tr key={i}>
                                    <td>{name}</td>
                                    <td>{tier}</td>
                                    <td>{count}</td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>

                <div>
                    <h3>Total Dust:</h3>
                    <table>
                        <tbody>
                            <tr>
                                <th>Dust Tier</th>
                                <th>Amount</th>
                            </tr>
                            {Object.keys(totalDust).map((tier, i) =>
                                <tr key={i}>
                                    <td>{tier}</td>
                                    <td>{totalDust[tier]}</td>
                                </tr>
                            )}
                        </tbody>
                    </table>
                </div>
            </div>
        }
    </div>

}

export default Table