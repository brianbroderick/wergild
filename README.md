Wergild is a MUD written in Go.

## Adding commands ##

To add commands: 
First add a token, and possible aliases in token.go. We'll use LOOK as our example. LOOK has an alias L. LOOK gets added to the token constants, and L gets added to convertToTokens. You'll also need to add an entry for both constants in the tokens map. This is so we can look up a token via its string representation.

Next, add a handler in parse_tree.go like this:  

```
Language.Handle(LOOK, func(p *Parser) (Statement, error) {
		return p.parseLookStatement() // resolver method
	})
```

`LOOK` matches your token. You won't need to do this for the alias because it's already been converted. 

Next, create a struct that will use the Statement interface, along with the appropriate methods. Currently this is String(), setPlayer(), and execute(). Then add a resolver like parseLookStatement() in the example above.

Finally, have fun using your new command to interact with the world.

One of the simplest commands is QUIT, which you can use as a template. It's located in stmt_admin.go. One of the commands with the most scanner variations is LOOK, which is located in stmt_senses.go. 

## Current commands ##

### LOOK ###

`look` alias `l`, allows you to look at the description of a room. You can also `l at <interesting thing>`. If the room has a description for it, it'll show you a description of it. For example, `l at chair`. You can also look in a direction, but it doesn't do much more than print a message at this time. For example, `l n` i.e. `look north`.

### FEELINGS ###

A feeling is like a text representation of an emoji. Type `laugh` and you will display that `You fall down laughing`. This will become a way to express feelings to other players. 

### DIRECTIONS ###

Without the ability to move from one room to another, this would be a pretty boring MUD. Typing a direction like `north` (or the alias `n`) will allow you to move in that direction, if there's an exit available to you for that direction. There are also aliases set up for the first letter of each direction. Currently implemented directions are north, south, east, west, up, and down. 

### REPEATED COMMANDS ###

By starting your command with a number, it will execute that command that many times. It will err above 50. For example `10 look`. Be aware that causing unnecessary spam on the line may cause other players to get pissed and potentially retaliate. 