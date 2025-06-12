package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

// Built-in quotes that are used when no external file is specified.
var builtInQuotes = []string{
	"I would rather have questions that can't be answered than answers that can't be questioned.\" -- Richard Feynman",
	"Who controls the past controls the future; who controls the present controls the past. -- George Orwell, \"1984\"",
	"When an immigrant resides with you in your land, you shall not oppress the immigrant. The immigrant who resides with you shall be to you as the citizen among you; you shall love the immigrant as yourself, for you were immigrants in the land of Egypt: I am the Lord your God. -- Leviticus 19:33-34",
	"Once, men turned their thinking over to machines in the hope that this would set them free.  But that only permitted other men with machines to enslave them. -- Frank Herbert, \"Dune\"",
	"If my words did glow with the gold of sunshine / And my tunes were played on the harp unstrung / Would you hear my voice come through the music? / Would you hold it near, as it were your own? -- Robert Hunter, \"Ripple\"",
	"In the land of the dark, the ship of the sun is drawn by the Grateful Dead. -- The Egyptian Book of the Dead",
	"The illegality of cannabis is outrageous, an impediment to full utilization of a drug which helps produce the serenity and insight, sensitivity and fellowship so desperately needed in this increasingly mad and dangerous world. -- Carl Sagan",
	"I before E, except when your foreign neighbor Keith receives eight counterfeit beige sleighs from feisty caffeinated weightlifters.  Wierd. -- Unknown",
	"Again I tell you, it is easier for a camel to go through the eye of a needle than for a rich man to enter the kingdom of God. -- Matthew, 19:24",
	"Let me never fall into the vulgar mistake of dreaming that I am persecuted whenever I am contradicted. -- Ralph Waldo Emerson",
	"Heaven is where the police are British, the lovers French, the mechanics German, the chefs Italian, and it is all organized by the Swiss.  Hell is where the police are German, the lovers Swiss, the mechanics French, the chefs British, and it is all organized by the Italians. -- Unknown",
	"In the world I see - you are stalking elk through the damp canyon forests around the ruins of Rockefeller Center.  You'll wear leather clothes that will last you the rest of your life.  You'll climb the wrist-thick kudzu vines that wrap the Sears Tower.  And when you look down, you'll see tiny figures pounding corn, laying strips of venison on the empty car pool lane of some abandoned superhighway. -- Tyler Durden",
	"It is difficult to get a man to understand something, when his salary depends upon his not understanding it! -- Upton Sinclair",
	"When a man sells his daughter as a slave, she will not be freed at the end of six years as the men are.  If she does not please the man who bought her, he may allow her to be bought back again.  But he is not allowed to sell her to foreigners, since he is the one who broke the contract with her.  And if the slave girl's owner arranges for her to marry his son, he may no longer treat her as a slave girl, but he must treat her as his daughter.  If he himself marries her and then takes another wife, he may not reduce her food or clothing or fail to sleep with her as his wife.  If he fails in any of these three ways, she may leave as a free woman without making any payment. -- Exodus 21:7-11",
	"Many that live deserve death.  And some die that deserve life.  Can you give it to them?  Then be not too eager to deal out death in the name of justice, fearing for your own safety.  Even the wise cannot see all ends. -- Gandalf, The Fellowship of The Ring",
	"I expect death to be nothingness and, for removing me from all possible fears of death, I am thankful to atheism. -- Isaac Asimov",
	"Those who cannot remember the past are condemned to repeat it. -- George Santayana",
	"It has been well said that if triangles had a god, they would give him three sides. -- Charles-Louis Montesquieu, Persian Letters, 1721",
	"On voting: http://www.youtube.com/watch?v=6dqsNrmXgP0 -- George Carlin",
	"That which can be asserted without evidence can be dismissed without evidence. -- Christopher Hitchins",
	"Ignorance is preferable to error; and he is less remote from the truth who believes nothing, than he who believes what is wrong. -- Thomas Jefferson",
	"Is God willing to prevent evil but not able?  Then he is not omnipotent.  Is he able but not willing?  Then he is malevolent.  Is he both able and willing?  Then whence cometh evil?  Is he neither able nor willing?  Then why call him God. -- Epicurus (341 BC - 270 BC)",
	"The gambling known as business looks with austere disfavor upon the business known as gambling. -- Ambrose Bierce",
	"Experience is the name everyone gives to their mistakes. -- Oscar Wilde",
	"As a matter of principle, I never attend the first annual anything. -- George Carlin",
	"It is easier to forgive an enemy than to forgive a friend. -- William Blake",
	"Fallacies do not cease to be fallacies because they become fashions. -- G. K. Chesterton",
	"College isn't the place to go for ideas. -- Helen Keller",
	"In great affairs men show themselves as they wish to be seen; in small things they show themselves as they are. -- Nicolas Chamfort",
	"The Moving Finger writes; and, having writ, / Moves on: nor all thy Piety nor Wit / Shall lure it back to cancel half a Line, / Nor all thy Tears wash out a Word of it. -- Omar Khayyám",
	"Everything good that happened to me happened by accident.  I was not filled with ambition nor fired by a drive toward a clear-cut goal.  I never knew exactly where I was going. -- Jack Benny",
	"There is danger from all men.  The only maxim of a free government ought to be to trust no man living with power to endanger the public liberty. -- John Adams, 1772",
	"I would rather be exposed to the inconveniences attending too much liberty than to those attending too small a degree of it. -- Thomas Jefferson (1743 - 1826)",
	"This last temptation is the greatest treason: to do the right deed for the wrong reason. -- T. S. Eliot",
	"An adventure is only an inconvenience rightly considered. -- G. K. Chesterton",
	"All you have to decide is what you're going to do with the time you are given. -- Gandalf (J.R.R. Tolkein)",
	"It is the mark of an educated mind to be able to entertain a thought without accepting it. -- Aristotle",
	"If you give someone Fortran, he has Fortran. If you give someone Lisp, he has any language he pleases. -- Guy L. Steele ",
	"When you have eliminated the impossible, whatever remains, however improbable, must be the truth. -- Sherlock Holmes (Sir Arthur Conan Doyle)",
	"There are three principal ways to lose money: wine, women, and engineers. While the first two are more pleasant, the third is by far the more certain. -- Baron Rothschild",
	"Doubt is not a pleasant condition, but certainty is an absurd one. -- Voltaire",
	"Outside of a dog, a book is man's best friend.  Inside of a dog, it's too dark to read. -- Groucho Marx",
	"Those who can make you believe absurdities can make you commit atrocities. -- Voltaire",
	"Behind it all is surely an idea so simple, so beautiful, that when we grasp it - in a decade, a century, or a millennium - we will all say to each other, how could it have been otherwise? How could we have been so stupid for so long? -- John Archibald Wheeler",
	"Peace cannot be kept by force.  It can only be achieved by understanding. -- Albert Einstein",
	"Good breeding consists in concealing how much we think of ourselves and how little we think of the other person. -- Mark Twain",
	"When a noble life has prepared old age, it is not decline that it reveals, but the first days of immortality. -- Germaine De Stael",
	"Without humility there can be no humanity. -- Sir John Buchan",
	"Any sufficiently advanced technology is indistinguishable from magic. -- Sir Arthur C. Clarke",
	"People sleep peaceably in their beds at night only because rough men stand ready to do violence on their behalf. -- George Orwell",
	"Education is the ability to listen to almost anything without losing your temper or your self-confidence. -- Robert Frost",
	"No opera plot can be sensible, for people do not sing when they are feeling sensible. -- W. H. Auden",
	"We confess our little faults to persuade people that we have no large ones. -- Francois de La Rochefoucauld",
	"The visionary lies to himself, the liar only to others. -- Friedrich Nietzsche",
	"Who is rich?  He that is content.  Who is that?  Nobody. -- Benjamin Franklin",
	"An author is a fool who, not content with boring those he lives with, insists on boring future generations. -- Charles de Montesquieu",
	"Few things are harder to put up with than the annoyance of a good example. -- Mark Twain",
	"I can forgive Alfred Nobel for having invented dynamite, but only a fiend in human form could have invented the Nobel Prize. -- George Bernard Shaw",
	"My wife and I were happy for twenty years.  Then we met. -- Rodney Dangerfield",
	"Football is a mistake.  It combines the two worst elements of American life.  Violence and committee meetings. -- George F. Will",
	"If all else fails, immortality can always be assured by spectacular error. -- John Kenneth Galbraith",
	"The height of cleverness is to be able to conceal it. -- Francois de La Rochefoucauld",
	"Those are my principles, and if you don't like them ... well, I have others. -- Groucho Marx",
	"It's a dangerous business going out your front door. -- J. R. R. Tolkien",
	"A single death is a tragedy; a million deaths is a statistic. -- Joseph Stalin",
	"Life is full of misery, loneliness, and suffering -- and it's all over much too soon. -- Woody Allen",
	"In mathematics you don't understand things.  You just get used to them. -- Johann von Neumann",
	"A government that robs Peter to pay Paul can always depend on the support of Paul. -- George Bernard Shaw",
	"Good judgement comes from experience.  Experience comes from bad judgement.  -- Unknown",
	"A foolish consistency is the hobgoblin of little minds. -- Ralph Waldo Emerson",
	"It is difficult to get a man to understand something when his salary depends upon his not understanding it. -- Upton Sinclair",
	"I am not one of those who in expressing opinions confine themselves to facts. -- Mark Twain",
	"If you give someone a program, you will frustrate them for a day; if you teach them how to program, you will frustrate them for a lifetime. -- Anonymous",
	"A lie gets halfway around the world before the truth has a chance to get its pants on. -- Winston Churchill",
	"Power, like a desolating pestilence, Pollutes whate'er it touches ... -- Percy Bysshe Shelley",
	"The Undiscovered Country / from whose bourn no traveller returns. -- William Shakespeare",
	"The question of whether a computer can think is no more interesting than the question of whether a submarine can swim. -- Edsger W. Dijkstra",
	"A man more right than his neighbors already constitutes a majority of one. -- Henry David Thoreau",
	"A man can be what he wants, but he cannot want what he wants. -- T. E. Lawrence",
	"The time you enjoy wasting is not wasted time. -- Bertrand Russell",
	"Debugging is twice as hard as writing the code in the first place.  Therefore, if you write the code as cleverly as possible, you are, by definition, not smart enough to debug it. -- Brian W. Kernighan",
	"The great enemy of the truth is very often not the lie -- deliberate, contrived and dishonest -- but the myth -- persistent, persuasive, and unrealistic. -- John F. Kennedy, Commencement address, Yale University, June 11, 1962",
	"Human history becomes more and more a race between education and catastrophe. -- H. G. Wells",
	"If only there were evil people somewhere insidiously committing evil deeds, and it were necessary only to separate them from the rest of us and destroy them.  But the line dividing good and evil cuts through the heart of every human being, and who is willing to destroy his (or her) own heart? -- Alexander Solzhenitsyn",
	"The empires of the future are the empires of the mind. -- Winston Churchill",
	"He who speaks does not know. He who knows does not speak. -- Lao-Tzu, \"Tao Te Ching\"",
	"The difference between what the most and the least learned people know is inexpressibly trivial in relation to that which is unknown. -- Albert Einstein",
	"Those who dream by night in the dusty recesses of their minds awake to find that all was vanity; but the dreamers of day are dangerous men, that they may act their dreams with open eyes to make it possible. -- T. E. Lawrence (of Arabia)",
	"A man who does not read good books has no advantage over the man who can't read them. -- Mark Twain",
	"Whatever you can do, or dream you can, begin it.  Boldness has genius, power and magic in it. -- Goethe",
	"A hundred times every day I remind myself that my inner and outer life depend on the labors of other men, living and dead, and that I must exert myself in order to give in the same measure as I have received. -- Albert Einstein",
	"We must not cease from exploration.  And the end of all our exploring will be to arrive where we began and to know the place for the first time. -- T. S. Elliot",
	"In theory there is no difference between theory and practice.  In practice, there is. -- Yogi Berra",
	"Under a government which imprisons any unjustly, the true place for a just man is in prison. -- Henry David Thoreau (1817 - 1862)",
	"I am returning this otherwise good typing paper to you because someone has printed gibberish all over it and put your name at the top. -- An English Professor, Ohio University",
	"A common mistake people make when trying to design something completely foolproof is to underestimate the ingenuity of complete fools. -- Douglas Adams (1952 - 2001), \"Mostly Harmless\"",
	"Don't let it end like this. Tell them I said something. -- Pancho Villa (1877 - 1923), last words",
	"It was a bright cold day in April, and the clocks were striking Thirteen. -- George Orwell, \"1984\"",
	"Facts do not cease to exist just because they are ignored. -- Aldous Huxley (1894 - 1963)",
	"We Americans live in a nation where the medical-care system is second to none in the world, unless you count maybe 25 or 30 little scuzzball countries like Scotland that we could vaporize in seconds if we felt like it. -- Dave Barry",
	"It is easier to port a shell than a shell script. -- Larry Wall",
	"There are seven sins in the world: wealth without work, pleasure without conscience, knowledge without character, commerce without morality, science without humanity, worship without sacrifice and politics without principle. -- Mahatma Gandhi (1869 - 1948)",
	"Happy is he who gets to know the reasons for things. -- Virgil (70 BC - 19 BC)",
	"The future is a hundred thousand threads, but the past is a fabric that can never be rewoven. -- Orson Scott Card, Xenocide",
	"What we obtain too cheap, we esteem too lightly. It is dearness only that gives everything its value. -- Thomas Paine (1737 - 1809)",
	"It requires a very unusual mind to undertake the analysis of the obvious. -- Alfred North Whitehead (1861 - 1947)",
	"Power is not revealed by striking hard or often, but by striking true. -- Honore de Balzac",
	"The Bible tells us to love our neighbors, and also to love our enemies; probably because generally they are the same people. -- G. K. Chesterton",
	"One repays a teacher badly if one only remains a pupil. -- Nietzsche",
	"When you have an efficient government, you have a dictatorship. -- Harry S Truman (1884 - 1972)",
	"Although the whole of this life were said to be nothing but a dream and the physical world nothing but a phantasm, I should call this dream or phatasm real enough, if, using reason well, we were never deceived by it. -- Baron Gottfried Wilhelm von Leibniz",
	"It is well, I die hard, but I am not afraid to go. -- George Washington (1732 - 1799), last words, 14 December 1799.",
	"Don't waste yourself in rejection, nor bark against the bad, but chant the beauty of the good. -- Ralph Waldo Emerson (1803 - 1882)",
	"I said to the almond tree, \"Sister, speak to me of God.\", and the almond tree blossomed. -- N. Kazantzakis",
	"There is a theory which states that if ever anybody discovers exactly what the Universe is for and why it is here, it will instantly disappear and be replaced by something even more bizarre and inexplicable.  There is another theory which states that this has already happened. -- Douglas Adams (1952 - 2001)",
	"Men have become the tools of their tools. -- Henry David Thoreau (1817 - 1862)",
	"I know not with what weapons World War III will be fought, but World War IV will be fought with sticks and stones. -- Albert Einstein (1879 - 1955)",
	"There are two kinds of drivers: maniacs and idiots.  The maniacs are everyone going faster than you, and the idiots are everyone going slower than you. -- George Carlin",
	"We're all in the gutter, but some of us are looking at the stars. -- Oscar Wilde (1854 - 1900)",
	"Ignorance, the root and the stem of every evil. -- Plato",
	"When armies are mobilized and issues are joined, the man who is sorry over the fact will win. -- Lao-tzu, The Way of Lao-tzu",
	"The only thing that interferes with my learning is my education. -- Albert Einstein",
	"I use emotion for the many and reserve reason for the few. -- Adolph Hitler",
	"Clothes make the man.  Naked people have little or no influence on society. -- Mark Twain",
	"Do not meddle in the affairs of wizards, for they are subtle and quick to anger. -- J. R. R. Tolkien",
	"Isn't it strange that I who have written only unpopular books should be such a popular fellow? -- Albert Einstein (1879 - 1955)",
	"Why do we drive on parkways and park on driveways?",
	"What luck for rulers that men do not think. -- Adolf Hitler",
	"The human brain starts working the moment you are born and never stops until you stand up to speak in public. -- George Jessel",
	"He is a hard man who is only just, and a sad one who is only wise. -- Voltaire (1694 - 1778)",
	"Observe your enemies, for they first find out your faults. -- Antisthenes (445 BC - 365 BC)",
	"Horse sense is the thing a horse has which keeps it from betting on people. -- W. C. Fields",
	"I must be cruel only to be kind; Thus bad begins, and worse remains behind. -- William Shakespeare",
	"What I look forward to is continued immaturity followed by death. -- Dave Barry",
	"The optimist proclaims that we live in the best of all possible worlds; and the pessimist fears this is true. -- James Branch Cabell (1879 - 1958)",
	"He that would make his own liberty secure must guard even his enemy from oppression; for if he violates this duty he establishes a precedent that will reach to himself. -- Thomas Paine (1737 - 1809)",
	"If God did not exist, it would be necessary to invent him. -- Voltaire",
	"Martyrdom is the only way in which a man can become famous without ability. -- George Bernard Shaw",
	"Blessed is the man who, having nothing to say, abstains from giving wordy evidence of the fact. -- George Eliot",
	"There's a fine line between fishing and just standing on the shore like an idiot. -- Steven Wright",
	"Three may keep a secret, if two of them are dead. -- Benjamin Franklin",
	"The more laws and order are made prominent, the more thieves and robbers there will be. -- Lao-tzu, The Way of Lao-tzu",
	"To someone seeking power, the poorest man is the most useful. -- Sallust (86 BC - 34 BC)",
	"If a 'religion' is defined to be a system of ideas that contains unprovable statements, then Godel taught us that mathematics is not only a religion, it is the only religion that can prove itself to be one. -- John Barrow",
	"I'm a great believer in luck, and I find the harder I work the more I have of it. -- Thomas Jefferson",
	"Nearly all men can stand adversity, but if you want to test a man's character, give him power. -- Abraham Lincoln",
	"Tragedy is when I cut my finger.  Comedy is when you walk into an open sewer and die. -- Mel Brooks",
	"I wasted time, and now doth time waste me. -- William Shakespeare",
	"Whatever you do will be insignificant, but it is very important that you do it. -- Mahatma Gandhi",
	"Under capitalism, man exploits man.  Under communism, it's just the opposite. -- John Kenneth Galbraith",
	"Corporations have been enthroned, an era of corruption in high places will follow, and the money-power of the country will endeavor to prolong its reign by working upon the prejudices of the people until the wealth is aggregated in a few hands and the Republic is destroyed. -- Abraham Lincoln (1809 - 1865), quoted in Jack London's \"The Iron Heel\"",
	"Reality is merely an illusion, albeit a very persistent one. -- Albert Einstein",
	"With or without religion, you would have good people doing good things and evil people doing evil things.  But for good people to do evil things, that takes religion. -- Steven Weinberg",
	"I believe that a scientist looking at nonscientific problems is just as dumb as the next guy. -- Richard Feynman",
	"Power tends to corrupt and absolute power corrupts absolutely. -- Lord Acton, in a letter to Bishop Mandell Creighton, April 3, 1887",
	"In science one tries to tell people, in such a way as to be understood by everyone, something that no one ever knew before.  But in poetry, it's the exact opposite. -- Paul Dirac (1902 - 1984)",
	"No place affords a more striking conviction of the vanity of human hopes than a public library. -- Samuel Johnson (1709 - 1784)",
	"As a well-spent day brings happy sleep, so life well used brings happy death. -- Leonardo da Vinci",
	"War is cruel and you cannot refine it. -- William Tecumseh Sherman",
	"When the rich make war it's the poor that die. -- Jean-Paul Sartre",
	"All that is gold does not glitter; not all those that wander are lost. -- J. R. R. Tolkien",
	"When a nation goes down, or a society perishes, one condition may always be found; they forgot where they came from.  They lost sight of what had brought them along. -- Carl Sandburg (1878 - 1967)",
	"The more I study religions the more I am convinced that man never worshipped anything but himself. -- Sir Richard Francis Burton",
	"I would therefore like to posit that computing's central challenge, viz. \"How not to make a mess of it,\" has /not/ been met. -- Edsger Dijkstra, 1930-2002",
	"If a person with multiple personalities threatens suicide, is that considered a hostage situation? -- George Carlin",
	"Indeed, it has been said that democracy is the worst form of government except all those other forms that have been tried from time to time. -- Sir Winston Churchill",
	"No human thing is of serious importance. -- Plato",
	"Those who profess to favor freedom, and yet deprecate agitation, are men who want rain without thunder and lightning.  They want the ocean without the roar of its many waters. -- Frederick Douglass",
	"The opposite of a correct statement is a false statement.  But the opposite of a profound truth may well be another profound truth. -- Niels Bohr (1885 - 1962)",
	"Whenever a man does a thoroughly stupid thing, it is always from the noblest motives. -- Oscar Wilde (1854 - 1900), The Picture of Dorian Gray, 1891",
	"Everything you can imagine is real. -- Pablo Picasso",
	"Passionate hatred can give meaning and purpose to an empty life. -- Eric Hoffer (1902 - 1983)",
	"Those who are too smart to engage in politics are punished by being governed by those who are dumber. -- Plato (427 BC - 347 BC)",
	"Do not fear death so much, but rather the inadequate life. -- Bertolt Brecht",
	"Pale Death with impartial tread beats at the poor man's cottage door and at the palaces of kings. -- Horace (65 BC - 8 BC), Odes",
	"All progress is based upon a universal innate desire on the part of every organism to live beyond its income. -- Samuel Butler",
	"It is not the man who has too little, but the man who craves more, that is poor. -- Seneca (5 BC - 65), Epistles",
	"Few things are impossible to diligence and skill.  Great works are performed not by strength, but perseverance. -- Samuel Johnson",
	"If you shoot at mimes, should you use a silencer? -- Steven Wright",
	"My religion consists of a humble admiration of the illimitable superior spirit who reveals himself in the slight details we are able to perceive with our frail and feeble mind. -- Albert Einstein",
	"Love is a snowmobile racing across the tundra and then suddenly it flips over, pinning you underneath.  At night, the ice weasels come. -- Matt Groening",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. -- Ralph Waldo Emerson",
	"You'll never have a quiet world till you knock the patriotism out of the human race. -- George Bernard Shaw, \"Misalliance\"",
	"Rather than love, than money, than fame, give me truth.  I sat at a table where were rich food and wine in abundance, and obsequious attendance, but sincerity and truth were not; and I went away hungry from the inhospitable board. -- Henry David Thoreau (1817 - 1862), Walden",
	"Why don't you ever see the headline \"Psychic Wins Lottery\"?",
	"The trouble with being punctual is that nobody's there to appreciate it. -- Franklin P. Jones",
	"Generosity is giving more than you can, and pride is taking less than you need. -- Kahlil Gibran",
	"It is impossible to enjoy idling thoroughly unless one has plenty of work to do. -- Jerome K. Jerome",
	"Nature is by and large to be found out of doors, a location where, it cannot be argued, there are never enough comfortable chairs. -- Fran Lebowitz",
	"Quality has to be caused, not controlled. -- Philip Crosby, Reflections on Quality",
	"Almost all absurdity of conduct arises from the imitation of those whom we cannot resemble. -- Samuel Johnson",
	"Niklaus Wirth has lamented that, whereas Europeans pronounce his name correctly (Ni-klows Virt), Americans invariably mangle it into (Nick-les Worth).  Which is to say that Europeans call him by name, but Americans call him by value.",
	"Channeling is just bad ventriloquism.  You use another voice, but people can see your lips moving. -- Penn Jillette",
	"If you pick up a starving dog and make him prosperous, he will not bite you; that is the principal difference between a dog and a man. -- Mark Twain (1835 - 1910)",
	"It is better to be violent, if there is violence in our hearts, than to put on the cloak of nonviolence to cover impotence. -- Mahatma Gandhi (1869 - 1948)",
	"All my possessions for a moment of time. -- Last words of Queen Elizabeth",
	"If I were to wish for anything, I should not wish for wealth and power, but for the passionate sense of potential -- for the eye which, ever young and ardent, sees the possible.  Pleasure disappoints; possibility never. -- Søren Kierkegaard",
	"I hear and I forget.  I see and I remember.  I do and I understand. -- Confucius",
	"Shouldn't the Psychic Friends call you instead?",
	"The weak can never forgive.  Forgiveness is the attribute of the strong. -- Mahatma Gandhi",
	"The basis for optimism is sheer terror. -- Oscar Wilde (1854 - 1900)",
	"It may be that our role on this planet is not to worship God -- but to create him. -- Arthur C. Clarke",
	"Patriotism is your conviction that this country is superior to all other countries because you were born in it. -- George Bernard Shaw",
	"Security is mostly superstition.  It does not exist in nature, nor do the children of men as a whole experience it.  Avoiding danger is no safer in the long run than outright exposure.  Life is either a daring adventure or nothing. -- Hellen Keller, \"The Open Door\", 1957",
	"Anyone who considers arithmetical methods of producing random digits is, of course, in a state of sin. -- John von Neumann",
	"Do radioactive cats have eighteen half-lives?",
	"Our thoughts are free. -- Cicero (106 BC - 43 BC)",
	"But in all my experience, I have never been in any accident ... or any sort worth speaking about.  I have seen but one vessel in distress in all my years at sea.  I never saw a wreck and never have been wrecked nor was I ever in any predicament that threatened to end in disaster of any sort. -- E. J. Smith, 1907, Captain, RMS Titanic",
	"Experience is that marvelous thing that enables you to recognize a mistake when you make it again. -- Franklin P. Jones",
	"Insanity: doing the same thing over and over again and expecting different results. -- Albert Einstein (1879 - 1955)",
	"Lose your dreams and you will loose your mind. -- Mick Jagger, \"Goodbye Ruby Tuesday\""}

// readQuotesFromFile reads quotes from the specified file and returns them as a slice.
func readQuotesFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var quotes []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Only add non-empty lines to avoid selecting blank lines.
		if len(line) > 0 {
			quotes = append(quotes, line)
		}
	}

	// Check for scanning errors.
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return quotes, nil
}

// showHelp displays usage information and exits.
func showHelp() {
	fmt.Println("quotebot - A random quote generator")
	fmt.Println()
	fmt.Println("USAGE:")
	fmt.Println("  quotebot [OPTIONS] [FILE]")
	fmt.Println()
	fmt.Println("DESCRIPTION:")
	fmt.Println("  Displays a randomly selected quote. By default, uses built-in quotes.")
	fmt.Println("  If a file is specified, quotes are read from that file instead.")
	fmt.Println("  Each quote in the file should be on a separate line.")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("  -h, --help    Show this help message and exit")
	fmt.Println()
	fmt.Println("ARGUMENTS:")
	fmt.Println("  FILE          Optional path to a file containing quotes")
	fmt.Println()
	fmt.Println("EXAMPLES:")
	fmt.Println("  quotebot                 # Use built-in quotes")
	fmt.Println("  quotebot quotes.txt      # Use quotes from quotes.txt")
	fmt.Println("  quotebot -h              # Show this help")
	os.Exit(0)
}

func main() {
	var quotes []string
	var err error

	// Check for help flags first.
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		showHelp()
	}

	// Check if a file path was provided as a command-line argument.
	if len(os.Args) > 1 {
		// Use quotes from the specified file.
		filename := os.Args[1]
		quotes, err = readQuotesFromFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading quotes from %s: %v\n", filename, err)
			os.Exit(1)
		}
	} else {
		// Use built-in quotes.
		quotes = builtInQuotes
	}

	// Ensure we have at least one quote to select from.
	if len(quotes) == 0 {
		fmt.Fprintf(os.Stderr, "No quotes available\n")
		os.Exit(1)
	}

	// Select a random quote and print it to standard output.
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
