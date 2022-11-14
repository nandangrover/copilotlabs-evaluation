class MultinomialNB {
    constructor(articles_per_tag) {
        this.alpha = 1;
        this.priors_per_tag = {};
        this.likelihood_per_word_per_tag = {};
        this.articles_per_tag = articles_per_tag;
        this.tags = Object.keys(articles_per_tag);
        this.train();
    }

    train() {
        let tag_counts_map = {};
        for (let tag of this.tags) {
            tag_counts_map[tag] = this.articles_per_tag[tag].length;
        }
        for (let tag of this.tags) {
            this.priors_per_tag[tag] = tag_counts_map[tag] / Object.values(tag_counts_map).reduce((a, b) => a + b, 0);
        }
        this.likelihood_per_word_per_tag = this.__get_word_likelihoods_per_tag();
    }

    predict(article) {
        let posteriors_per_tag = {};
        for (let tag in this.priors_per_tag) {
            posteriors_per_tag[tag] = Math.log(this.priors_per_tag[tag]);
        }
        for (let word of article) {
            for (let tag of this.tags) {
                posteriors_per_tag[tag] = posteriors_per_tag[tag] + Math.log(this.likelihood_per_word_per_tag[word][tag]);
            }
        }
        return posteriors_per_tag;
    }

    __get_word_likelihoods_per_tag() {
        let word_frequencies_per_tag = {};
        for (let tag of this.tags) {
            word_frequencies_per_tag[tag] = 0;
        }
        let total_word_count_per_tag = {};
        for (let tag of this.tags) {
            for (let article of this.articles_per_tag[tag]) {
                for (let word of article) {
                    word_frequencies_per_tag[word][tag] += 1;
                    total_word_count_per_tag[tag] += 1;
                }
            }
        }
        let word_likelihoods_per_tag = {};
        for (let tag of this.tags) {
            word_likelihoods_per_tag[tag] = 0.5;
        }
        for (let word in word_frequencies_per_tag) {
            for (let tag in word_frequencies_per_tag[word]) {
                word_likelihoods_per_tag[word][tag] = (word_frequencies_per_tag[word][tag] + 1 * this.alpha) / (
                    total_word_count_per_tag[tag] + 2 * this.alpha
                );
            }
        }
        return word_likelihoods_per_tag;
    }
}

let examples = {
    "politics": [[
        "trump impeachment",
        "hillary clinton",
        "barack obama",
        "donald trump",
        "hillary clinton",
        "barack obama",
        "donald trump",
        "hillary clinton",
    ]],
    "technology": [[
        "artificial intelligence",
        "machine learning",
        "deep learning",
        "computer science",
        "artificial intelligence",
        "machine learning",
        "deep learning",
        "computer science",
    ]],
    "science": [[
        "global warming",
        "climate change",
        "black hole",
        "dark matter",
        "global warming",
        "climate change",
        "black hole",
        "dark matter",
    ]],
}

let multinomial_nb = new MultinomialNB(examples);
let article = [
    "donald trump",
    "hillary clinton",
    "barack obama",
    "donald trump",
    "hillary clinton",
];
let expected = {'politics': -91.3016, 'sports': -87.1427, 'tech': -85.1920};
console.log(multinomial_nb.predict(article));