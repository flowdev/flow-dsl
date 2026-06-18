# Convert a Simple Flow Diagram DSL to Code and Documentation

Read a DSL for flows and convert it to documentation (MarkDown containing SVG images) or source code.

## FAQ

#### Why Use Flows For Solving Business Problems?

- Business people think in terms of (work) flows. \
  So this helps a lot to remove the friction between problem and solution space (business side and technical side). \
  It enables product people to always be on the same page with the engineers.
- No more: 'That requirement simply doesn't fit our design (or architecture).' \
  A small change in the problem domain always results in a small change in the solution domain
  since the solution is structured like the problem. \
  The solution can never be less complex than the problem it solves but it can easily be much more complex.
  This approach avoids this unnecessary complexity for all problems where it's a good fit.

#### What Technical Advantages Do Flows Have?

- They nest nicely. A flow can be a small part in another flow.
  This allows the construction of arbitrarily complex flows that can be easily followed in the documentation.
- Parts of a flow are independent of each other and they don't know about each other.
  This minimizes the "blast radius" of changes.
  In fact many changes can be done without any change to any existing component at all
  by simply adding a new component.

#### When Shouldn't I Use Flows?

In short: Whenever they aren't a good fit.

Flows usually fit business domains very well.
They can also be a good fit for some technical domains (e.g. sound engineering).

But they usually aren't a good fit for GUIs and other domains where an object oriented approach works much better.