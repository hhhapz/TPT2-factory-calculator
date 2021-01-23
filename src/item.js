import recipes from "./recipes"
import Part from "./part"

class Item {
    dust = 0

    constructor(name, tier, count = 1, index = "0", depth = 0) {
        this.name = name
        this.tier = tier
        this.count = count
        this.index = index
        this.depth = depth
        this.id = `${name}-${tier}-${index}-${depth}`
        this.components = []

        this.calc_components()
    }

    calc_components() {
        if (this.tier === 0) return [];

        const part = new Part(this.name, this.count)
        this.dust = part.totalDust()

        const recipe = recipes[this.tier][this.name]
        if (!recipe) return

        recipe.forEach((part, i) => {
            const item =
                new Item(
                    part.name,
                    part.tier === 0 ? this.tier : part.tier,
                    part.count * this.count,
                    `${this.index}${i}`,
                    this.depth + 1
                )
            this.components.push(item)
        })
    }

    tree() {
        const tree = { value: this.id, label: this.toString(), item: this }

        if (this.components.length > 0) {
            tree.children = []
        }
        this.components.forEach(item => {
            const [subTree] = item.tree()
            tree.children.push(subTree)
        })


        if (tree.children) tree.children.sort((a, b) => {
            if (a.item.tier === b.item.tier) return (a.item.name > b.item.name) ? -1 : 1
            return (a.item.tier > b.item.tier) ? -1 : 1
        })

        return [tree]
    }

    summary(parent = true, checked = []) {
        let items = {}

        if (checked.includes(this.id)) {
            if (parent) return []
            return {}
        }

        const key = `${this.name}@${this.tier}`
        if (this.components.length === 0) {
            items[key] = this.count
        } else this.components.forEach(part => {
            const summary = part.summary(false, checked)
            Object.keys(summary).forEach((key) => items[key] = (items[key] || 0) + summary[key])
        })

        if (!parent) return items

        return Object.keys(items)
            .map(key => {
                const [name, tier] = key.split("@")
                return [name, tier, items[key]]
            })
            .sort((a, b) => {
                if (a[1] === b[1]) return (a[0] > b[0]) ? -1 : 1
                return (a[1] > b[1]) ? -1 : 1
            })
    }

    totalDust(checked = []) {
        if (checked.includes(this.id)) return {}

        if (this.name === "Rubber") return { Rubber: this.count }

        const dust = {}

        if (this.components.length === 0) {
            dust[this.tier] = this.dust
            return dust
        }


        this.components.forEach((part) => {
            const subTotalDust = part.totalDust(checked)
            Object.keys(part.totalDust(checked)).forEach(subTier => {
                dust[subTier] = (dust[subTier] || 0) + subTotalDust[subTier]
            })
        })

        return dust
    }

    toString() {
        return `T${this.tier} ${this.name} - ${this.count}`
    }

    equals(other) {
        return this.tier === other.tier && this.name === other.name
    }
}

export default Item