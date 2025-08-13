# Product Documentation

## Background

Historically, The Linux Foundation has leveraged [a common
repository](https://github.com/linuxfoundation/lfx-product-documentation/tree/main)
to coalesce all the LFX product documentation into a single location. This was deployed
as [a static site](https://docs.linuxfoundation.org/lfx/).

## Single Repository Approach: Pros and Cons

### Pros of Centralized Documentation Repository

1. **Unified User Experience**: All product documentation is accessible from a
single location, providing a consistent navigation and search experience for
users.

2. **Cross-Product Discoverability**: Users can easily discover related features
across different products when documentation is centralized.

3. **Consistent Styling and Branding**: Single repository allows for uniform
styling, branding, and documentation standards across all products.

4. **Simplified Deployment**: One deployment pipeline for all documentation
updates, reducing infrastructure complexity.

5. **Centralized Governance**: Easier to enforce documentation standards, review
processes, and quality control from a single location.

6. **Reduced Duplication**: Easier to identify and eliminate duplicate content
across products.

7. **Unified Analytics**: Single source for documentation usage analytics and
user behavior insights.

### Cons of Centralized Documentation Repository

1. **Single Point of Failure**: Issues with the central repository can affect
documentation for all products simultaneously.

2. **Bottleneck in Updates**: All documentation changes must go through a single
review and deployment process, potentially slowing down updates.

3. **Dependency on Central Team**: Product teams become dependent on a central
technical writing team for documentation updates.

4. **Version Synchronization Challenges**: Difficulty in keeping documentation
synchronized with product releases when managed separately.

5. **Reduced Product Team Ownership**: Product teams have less direct control
over their documentation, potentially leading to outdated or inaccurate content.

6. **Scalability Issues**: As the number of products grows, the central
repository becomes harder to manage effectively.

7. **Knowledge Silos**: Central technical writers may not have deep expertise in
all products, leading to less accurate or comprehensive documentation.

## New Embedded Documentation Approach

### Transition Strategy

We are transitioning from a centralized documentation model to an **embedded
documentation approach** where product documentation lives alongside the product
code. This change addresses the limitations of the single repository model and
aligns with modern DevOps practices.

#### Success metrics and guardrails

- Target lead time from docs PR open → publish: 1 day.
- Freshness SLA: docs updated within 1 day of feature release.
- Rollback plan: if KPIs regress for two consecutive sprints, pause new
migrations and route publishing back through the central pipeline until issues
are remediated.

### Key Changes

1. **Elimination of Central Technical Writer Role**: Utilize a federated
documentation team approach.

2. **Product Manager Ownership**: Each product manager is now responsible for
authoring and maintaining their product's documentation.

3. **Embedded Documentation**: Documentation is stored within each product's
codebase alongside front-end code and other product assets.

4. **Integrated Development Workflow**: Documentation changes follow the same
development process as code changes.

### Embedded Documentation Benefits

1. **Improved Accuracy**: Documentation is maintained by those who know the
product best - the product managers and developers.

2. **Faster Updates**: Documentation can be updated and deployed alongside
product changes, ensuring synchronization.

3. **Better Ownership**: Product teams have direct control over their
documentation, leading to more comprehensive and up-to-date content.

4. **Reduced Dependencies**: No longer dependent on a central team for
documentation updates.

5. **Contextual Documentation**: Documentation can be more closely tied to
specific features and code components.

6. **Agile Documentation**: Documentation can evolve iteratively with product
development cycles.

7. **Documentation Links**: Documentation links are now embedded in the product
codebase, making it easier to find and navigate documentation for a given
product.

8. **Documentation Styling**: Documentation styling now follows the product's
branding and design guidelines. As products are updated, the documentation
styling will be updated to match.

9. **Documentation Search**: Documentation search is now integrated into the
product's search functionality.

10. **Documentation Versioning**: Documentation versioning is now integrated into the product's versioning system.

11. **Documentation Accessibility**: Documentation is now accessible to all
users, regardless of their location or device.

12. **Author Documentation**: Write and maintain comprehensive documentation for
their products. The primary responsibility of the product manager is to ensure
that the documentation is complete, accurate, and up-to-date. The developer is
responsible for authoring technical sections, examples, and API references.
Teams must adopt the shared docs theme and templates; packaging and deployment
are handled by a standardized docs CI pipeline integrated with product releases.
Teams must adopt the shared docs theme and templates; packaging and deployment
are handled by a standardized docs CI pipeline integrated with product releases.

13. **Author Documentation**: Write and maintain comprehensive documentation for
their products. The primary responsibility of the product manager is to ensure
that the documentation is complete, accurate, and up-to-date. The developer is
responsible for styling the documentation to match the product's branding and
design guidelines as well as packaging and deploying the documentation along
with the product releases.

14. **Submit Change Requests**: All documentation changes will follow the standard
process for submitting change requests.

15. **Review Process**: Documentation changes must go through the same review and
approval process as code changes. Feedback, comments, decisions, and approvals
are visible as part of the review process.

16. **Quality Assurance**: Ensure documentation meets established standards for clarity, completeness, and accuracy.

17. **Deployment Integration**: Documentation changes are deployed alongside product changes to production.

### Challenges and Considerations

1. **Documentation Standards**: Need to establish and maintain consistent
documentation standards across all product teams. The same documentation
standards should be used for all products.

2. **Training Requirements**: Product managers and developers may need training
on technical writing best practices. The same training should be provided for
all product teams.

3. **Cross-Product References**: Need strategies for handling documentation that
references multiple products. The same cross-product references should be used
for all products.

4. **Search and Discovery**: May need to implement cross-product search
capabilities for users. The same search and discovery capabilities should be
used for all products.

5. **Documentation Tooling**: Teams will need appropriate tools and templates
for creating and maintaining documentation.

6. **Quality Control**: Need processes to ensure documentation quality without
central oversight. The same quality control processes should be used for all
product teams.

### Migration Timeline

The transition to embedded documentation should be planned carefully to ensure:

- Minimal disruption to existing documentation
- Proper training and onboarding for product teams
- Establishment of new processes and standards
- Gradual migration of existing documentation to new locations
- User communication about changes to documentation access
- Decommissioning of the existing documentation site
