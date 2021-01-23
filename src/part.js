import { config, parts } from "./recipes";

class Part {
    constructor(name, amount) {
        this.name = name
        this.amount = amount
    }

    totalDust() {
        let count = 0
        const children = parts[this.name]

        if (!children) return count

        children.forEach(item => {
            if (item.name === "Dust") {
                count += item.count
                return
            }

            const subPart = new Part(item.name, item.count)
            count += subPart.totalDust()
        })

        return count * this.amount
    }
}

export default Part;